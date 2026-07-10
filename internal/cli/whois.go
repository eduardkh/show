package cli

import (
	"context"
	"fmt"
	"net/netip"
	"strings"
	"time"

	"github.com/eduardkh/show/internal/whois"
	"github.com/spf13/cobra"
)

type whoisOptions struct {
	host    string
	port    int
	raw     bool
	timeout time.Duration
}

var whoisOpts whoisOptions

var whoisCmd = &cobra.Command{
	Use: "whois IP", Short: "Look up registration and routing information for an IP address",
	Long:    "Look up an IP address using its authoritative Regional Internet Registry.\nThe default host, auto, asks IANA to select the registry. Use --host to override it.",
	Example: "  show ip whois 8.8.8.8\n  show ip whois --raw 8.8.8.8\n  show ip whois --host whois.arin.net 8.8.8.8",
	Args:    cobra.ExactArgs(1), RunE: runWhois,
}

func runWhois(cmd *cobra.Command, args []string) error {
	query := args[0]
	if _, err := netip.ParseAddr(query); err != nil {
		return fmt.Errorf("invalid IP address %q: %w", query, err)
	}
	client := whois.Client{Timeout: whoisOpts.timeout}
	host := strings.ToLower(strings.TrimSpace(whoisOpts.host))
	if host == whois.AutoHost {
		var err error
		host, err = whois.DiscoverHost(cmd.Context(), client, query)
		if err != nil {
			return err
		}
	}
	response, err := client.Query(cmd.Context(), host, whoisOpts.port, query)
	if err != nil {
		return fmt.Errorf("WHOIS lookup failed: %w", err)
	}
	if whoisOpts.raw {
		_, err = cmd.OutOrStdout().Write(response)
		return err
	}
	var summary whois.Summary
	if strings.EqualFold(host, whois.CymruHost) {
		summary, err = whois.ParseCymru(query, response)
		if err != nil {
			return fmt.Errorf("parse Team Cymru response: %w", err)
		}
	} else {
		summary = whois.Normalize(query, whois.ParseRecords(host, response))
		if cymru, cymruErr := queryCymru(cmd.Context(), client, query); cymruErr == nil {
			summary.MergeMissing(cymru)
		}
	}
	printWhoisSummary(cmd, summary)
	return nil
}

func queryCymru(ctx context.Context, client whois.Client, query string) (whois.Summary, error) {
	raw, err := client.Query(ctx, whois.CymruHost, 43, " -v "+query)
	if err != nil {
		return whois.Summary{}, err
	}
	return whois.ParseCymru(query, raw)
}

func printWhoisSummary(cmd *cobra.Command, s whois.Summary) {
	fields := []struct{ label, value string }{
		{"IP Address", s.Query}, {"Range", s.Range}, {"BGP Prefix", s.Prefix}, {"Network Name", s.Name},
		{"Organization", s.Organization}, {"Autonomous System Number", s.ASN}, {"Autonomous System Name", s.ASName},
		{"Country", s.Country}, {"Abuse Email", s.AbuseEmail}, {"Registry", s.Registry}, {"WHOIS Server", s.Server},
	}
	for _, field := range fields {
		if field.value != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", field.label, field.value)
		}
	}
}

func completeWhoisHost(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0, len(whois.KnownHosts))
	for _, known := range whois.KnownHosts {
		completions = append(completions, known.Host+"\t"+known.Registry+" — "+known.Description)
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	// Linux WHOIS uses -h for host. Define help explicitly without a shorthand
	// so Cobra does not reserve -h when it initializes the command.
	whoisCmd.Flags().Bool("help", false, "help for whois")
	whoisCmd.Flags().StringVarP(&whoisOpts.host, "host", "h", whois.AutoHost, "WHOIS server (auto selects the authoritative registry)")
	whoisCmd.Flags().IntVarP(&whoisOpts.port, "port", "p", 43, "WHOIS server port")
	whoisCmd.Flags().BoolVar(&whoisOpts.raw, "raw", false, "print the unmodified WHOIS response")
	whoisCmd.Flags().DurationVar(&whoisOpts.timeout, "timeout", 10*time.Second, "connection and read timeout")
	if err := whoisCmd.RegisterFlagCompletionFunc("host", completeWhoisHost); err != nil {
		panic(err)
	}
	ipCmd.AddCommand(whoisCmd)
}

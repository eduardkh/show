package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func startHttpServer(srv *http.Server) error {
	http.HandleFunc("/", requestHandler)
	fmt.Println("Starting HTTP server on port 8080...")
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	// Get timestamp
	timestamp := time.Now().Format("02/01/2006 15:04:05")

	// Get requesting host
	host := r.RemoteAddr

	// Prepare strings to hold request details
	details := fmt.Sprintf("HOST: %s TIMESTAMP: %s\n\n", host, timestamp)
	method := fmt.Sprintf("HTTP Method: %s URI: %s\n\n", r.Method, r.RequestURI)
	headers := "HEADERS:\n"
	for name, headerValues := range r.Header {
		for _, h := range headerValues {
			headers += fmt.Sprintf("%v: %v\n", name, h)
		}
	}
	headers += "\n"

	// Read the body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	body := fmt.Sprintf("BODY:\n%s\n", string(bodyBytes))

	// Combine all details
	response := details + method + headers + body

	// Print to console
	fmt.Println(response)

	// Write response back to client
	fmt.Fprint(w, response)
}

// webhookCmd represents the webhook command
var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Starts an HTTP server to log HTTP requests",
	Long: `Starts a simple HTTP server that logs the details of incoming HTTP requests,
including the method, headers, body, URI, requesting host, and timestamp.`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := &http.Server{Addr: ":8080"}

		go func() {
			if err := startHttpServer(srv); err != nil {
				fmt.Fprintln(os.Stderr, "Error starting server:", err)
				os.Exit(1)
			}
		}()

		// Setting up signal capturing
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		// Waiting for SIGINT (Ctrl+C)
		<-stop

		// Graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Fprintln(os.Stderr, "Error during shutdown:", err)
		}
		fmt.Println("HTTP server stopped")
	},
}

func init() {
	rootCmd.AddCommand(webhookCmd)
}

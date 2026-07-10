package whois

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

const defaultMaxResponseBytes int64 = 4 << 20

// Client performs WHOIS queries over TCP.
type Client struct {
	Timeout          time.Duration
	MaxResponseBytes int64
}

// Query sends query to a WHOIS server and returns its unmodified response.
func (c Client) Query(ctx context.Context, host string, port int, query string) ([]byte, error) {
	if strings.TrimSpace(host) == "" {
		return nil, fmt.Errorf("WHOIS host cannot be empty")
	}
	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("invalid WHOIS port %d", port)
	}
	timeout := c.Timeout
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	limit := c.MaxResponseBytes
	if limit <= 0 {
		limit = defaultMaxResponseBytes
	}
	dialer := net.Dialer{Timeout: timeout}
	conn, err := dialer.DialContext(ctx, "tcp", net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		return nil, fmt.Errorf("connect to %s: %w", host, err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return nil, fmt.Errorf("set deadline for %s: %w", host, err)
	}
	if _, err := io.WriteString(conn, strings.TrimRight(query, "\r\n")+"\r\n"); err != nil {
		return nil, fmt.Errorf("send query to %s: %w", host, err)
	}
	response, err := io.ReadAll(io.LimitReader(conn, limit+1))
	if err != nil {
		return nil, fmt.Errorf("read response from %s: %w", host, err)
	}
	if int64(len(response)) > limit {
		return nil, fmt.Errorf("response from %s exceeds %d bytes", host, limit)
	}
	if len(response) == 0 {
		return nil, fmt.Errorf("no data received from %s", host)
	}
	return response, nil
}

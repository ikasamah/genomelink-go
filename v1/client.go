package genomelink

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"runtime"

	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const (
	// defaultEndpoint is the default base URL for GENOME LINK API v1.
	defaultEndpoint = "https://genomelink.io/v1"

	// endpointEnvVar is the environment variable name that overrides the defaultEndpoint.
	endpointEnvVar = "GENOMELINK_V1_ENDPOINT"
)

var (
	version    = "0.0.1"
	projectURL = "https://github.com/ikasamah/genomelink-go"
	userAgent  = fmt.Sprintf("GENOMELINK-Go/%s (+%s; %s)", version, projectURL, runtime.Version())
)

// Client represents GENOMELINK API Client.
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Logger     *zap.Logger
}

// DefaultClient returns a client that connects to GENOME LINK API.
func DefaultClient(ctx context.Context, token string) *Client {
	endpoint := os.Getenv(endpointEnvVar)
	if endpoint == "" {
		endpoint = defaultEndpoint
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client, err := NewClient(endpoint, tc, nil)
	if err != nil {
		panic(err)
	}
	return client
}

// NewClient creates a new GENOME LINK client from given url, http client, and logger.
func NewClient(urlStr string, client *http.Client, logger *zap.Logger) (*Client, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if client == nil {
		client = http.DefaultClient
	}

	if logger == nil {
		logger = zap.NewNop()
	}

	return &Client{
		URL:        parsedURL,
		HTTPClient: client,
		Logger:     logger,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, query url.Values, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func (c *Client) decodeResponse(resp *http.Response, dest interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(dest)
}

func dumpRequest(req *http.Request) []byte {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		dump = []byte(fmt.Sprintf("%+v", req))
	}
	return dump
}

func dumpResponse(resp *http.Response) []byte {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		dump = []byte(fmt.Sprintf("%+v", resp))
	}
	return dump
}

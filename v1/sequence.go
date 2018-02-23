package genomelink

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Region represent region of the chromosome.
type Region struct {
	Chromosome Chromosome
	Start      int
	End        int
}

// String builds the string by prescribed format.
func (r *Region) String() string {
	return fmt.Sprintf("%s:%d-%d", r.Chromosome, r.Start, r.End)
}

// Report fetches sequence API for the given region, returns io.ReadCloser.
func (c *Client) Sequence(ctx context.Context, region Region) (io.ReadCloser, error) {
	const spath = "/genomes/sequence"
	query := url.Values{"region": {region.String()}}

	// TODO: validate min/max position
	req, err := c.newRequest(ctx, http.MethodGet, spath, query, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create new request")
	}

	c.Logger.Debug("Request", zap.ByteString("request", dumpRequest(req)))
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to send http request: %s", dumpRequest(req))
	}
	c.Logger.Debug("Response", zap.ByteString("response", dumpResponse(resp)))

	return resp.Body, nil
}

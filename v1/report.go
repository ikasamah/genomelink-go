package genomelink

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// Report represents the response of report API
	Report struct {
		Summary    Summary    `json:"summary"`
		Phenotype  Phenotype  `json:"phenotype"`
		Population Population `json:"population"`
		Scores     []Score    `json:"scores"`
	}
	// Summary represents the summary object in report response
	Summary struct {
		Text     string   `json:"text"`
		Score    int      `json:"score"`
		Warnings []string `json:"warnings"`
	}
	// Score represents the score object in report response
	Score struct {
		Score int    `json:"score"`
		Text  string `json:"text"`
	}
)

// Report fetches report API for the given phenotype and population
func (c *Client) Report(ctx context.Context, name PhenotypeName, population Population) (Report, error) {
	spath := fmt.Sprintf("/reports/%s/", name)
	query := url.Values{"population": {string(population)}}

	req, err := c.newRequest(ctx, http.MethodGet, spath, query, nil)
	if err != nil {
		return Report{}, errors.Wrapf(err, "failed to create new request")
	}

	c.Logger.Debug("Request", zap.ByteString("request", dumpRequest(req)))
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return Report{}, errors.Wrapf(err, "failed to send http request: %s", dumpRequest(req))
	}
	c.Logger.Debug("Response", zap.ByteString("response", dumpResponse(resp)))

	var report Report
	if err := c.decodeResponse(resp, &report); err != nil {
		return Report{}, errors.Wrapf(err, "failed to parse response: %s", dumpResponse(resp))
	}
	return report, nil
}

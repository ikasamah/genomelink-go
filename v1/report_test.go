package genomelink

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestClient_Report(t *testing.T) {
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("%s/reports/%s", testHost, PhenotypeNameGeneticEyeColor),
		httpmock.NewStringResponder(http.StatusOK, `{
  "summary": {
    "text": "Tend to not have brown eyes, slightly",
    "score": 1,
    "warnings": [
      "this is DEMO data because this user does not have genome data"
    ]
  },
  "phenotype": {
    "url_name": "eye-color",
    "display_name": "Eye color",
    "category": "trait"
  },
  "population": "european",
  "scores": [
    {
      "score": 0,
      "text": "Tend to not have brown eyes"
    },
    {
      "score": 1,
      "text": "Tend to not have brown eyes, slightly"
    },
    {
      "score": 2,
      "text": "Intermediate"
    },
    {
      "score": 3,
      "text": "Slight tendency for brown eyes"
    },
    {
      "score": 4,
      "text": "Stronger tendency for brown eyes"
    }
  ]
}`))

	client := testClient()
	report, err := client.Report(ctx, PhenotypeNameGeneticEyeColor, PopulationEuropean)

	if err != nil {
		t.Fatal(err)
	}

	if report.Summary.Text != "Tend to not have brown eyes, slightly" {
		t.Fatalf("unexpected %q", report.Summary.Text)
	}
	if report.Summary.Score != 1 {
		t.Fatalf("unexpected %q", report.Summary.Score)
	}
	if len(report.Scores) != 5 {
		t.Fatalf("unexpected length %q", report.Scores)
	}
}

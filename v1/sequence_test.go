package genomelink

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestClient_Sequence(t *testing.T) {
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	result := "CACTAAGCACACAGAGAATAATGTCTAGAAT"
	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("%s/genomes/sequence", testHost),
		httpmock.NewStringResponder(http.StatusOK, result))

	client := testClient()
	r, err := client.Sequence(ctx, Region{
		Chromosome: Chromosome1,
		Start:      1,
		End:        30,
	})

	if err != nil {
		t.Fatal(err)
	}

	raw, _ := ioutil.ReadAll(r)
	if string(raw) != result {
		t.Fatalf("unexpected %q", raw)
	}
}

package genomelink

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"os"
)

const testHost = "https://test"

func testClient() *Client {
	client, err := NewClient(testHost, http.DefaultClient, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func TestDefaultClient(t *testing.T) {
	client := DefaultClient(context.Background(), "test")

	if client.URL.String() != defaultEndpoint {
		t.Fatalf("expected %q to be %q", client.URL.String(), defaultEndpoint)
	}
}

func TestDefaultClient_overrideURL(t *testing.T) {
	defer os.Setenv(endpointEnvVar, os.Getenv(endpointEnvVar))
	anotherEndpoint := "http://localhost:8000"

	err := os.Setenv(endpointEnvVar, anotherEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	client := DefaultClient(context.Background(), "test")

	if client.URL.String() != anotherEndpoint {
		t.Fatalf("expected %q to be %q", client.URL.String(), anotherEndpoint)
	}
}

func TestNewClient(t *testing.T) {
	newEndpoint := "http://127.0.0.1:8000"
	client, err := NewClient(newEndpoint, http.DefaultClient, nil)
	if err != nil {
		t.Fatal(err)
	}

	if client.URL.String() != newEndpoint {
		t.Fatalf("expected %q to be %q", client.URL.String(), newEndpoint)
	}

	if client.HTTPClient != http.DefaultClient {
		t.Fatalf("given client is not used")
	}
}

func TestClient_newRequest(t *testing.T) {
	ctx := context.Background()
	client := testClient()

	req, err := client.newRequest(ctx, http.MethodGet, "/foo/bar", url.Values{"key": []string{"value"}}, nil)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("User-Agent") == "" {
		t.Fatalf("user-agent header is not set: %s", dumpRequest(req))
	}

	expected := testHost + "/foo/bar?key=value"
	if req.URL.String() != expected {
		t.Fatalf("expected %q to be %q", req.URL.String(), expected)
	}
}

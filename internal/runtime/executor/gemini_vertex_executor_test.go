package executor

import (
	"encoding/json"
	"testing"

	cliproxyauth "github.com/router-for-me/CLIProxyAPI/v6/sdk/cliproxy/auth"
)

func TestVertexCredsAcceptsAuthorizedUser(t *testing.T) {
	auth := &cliproxyauth.Auth{
		Metadata: map[string]any{
			"project_id": "cliproxyapp",
			"location":   "us-central1",
			"authorized_user": map[string]any{
				"type":          "authorized_user",
				"client_id":     "client-id",
				"client_secret": "client-secret",
				"refresh_token": "refresh-token",
			},
		},
	}

	projectID, location, credsJSON, err := vertexCreds(auth)
	if err != nil {
		t.Fatalf("vertexCreds returned error: %v", err)
	}
	if projectID != "cliproxyapp" {
		t.Fatalf("projectID = %q, want cliproxyapp", projectID)
	}
	if location != "us-central1" {
		t.Fatalf("location = %q, want us-central1", location)
	}
	var creds map[string]any
	if err := json.Unmarshal(credsJSON, &creds); err != nil {
		t.Fatalf("credentials json is invalid: %v", err)
	}
	if got := creds["type"]; got != "authorized_user" {
		t.Fatalf("credentials type = %v, want authorized_user", got)
	}
	if got := creds["refresh_token"]; got != "refresh-token" {
		t.Fatalf("refresh_token = %v, want refresh-token", got)
	}
}

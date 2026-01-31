package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyNoToken(t *testing.T) {
	header := make(http.Header)
	_, err := GetAPIKey(header)
	if err == nil {
		t.Errorf("no token provided -> should have failed")
	}
}

func TestGetAPIKeyEmptyToken(t *testing.T) {
	header := make(http.Header)
	header.Add("Authentication", "")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Errorf("empty token -> should have failed")
	}
}

func TestGetAPIKeyWrongTokenType(t *testing.T) {
	header := make(http.Header)
	header.Add("Authorization", "Bearer Token")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Errorf("wrong token type -> should have failed")
	}
}

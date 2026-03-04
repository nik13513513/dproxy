package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireBasicAuthUsesRealm(t *testing.T) {
	// preserve original value so other tests are not affected
	orig := Realm
	defer func() { Realm = orig }()

	Realm = "customrealm"
	ctx := context.Background()
	rr := httptest.NewRecorder()
	req := &http.Request{}

	user, ok := requireBasicAuth(ctx, rr, req, "", nil)
	if ok {
		t.Fatalf("expected ok=false when no credentials provided, got user=%q", user)
	}

	head := rr.Header().Get("Proxy-Authenticate")
	want := `Basic realm="customrealm"`
	if head != want {
		t.Errorf("unexpected Proxy-Authenticate header: got %q, want %q", head, want)
	}
	if rr.Code != 407 {
		t.Errorf("expected status 407, got %d", rr.Code)
	}
}

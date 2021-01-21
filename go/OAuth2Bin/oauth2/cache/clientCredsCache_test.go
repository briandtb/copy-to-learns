package cache

import "testing"

// TestClientCredsFlow tests the entirely of the functions set of authCodeStore
// as they would be used by the Implicit Grant flow
func TestClientCredsFlow(t *testing.T) {
	// Generating a token which would be done once the user authorizes
	// the client application
	token, err := NewClientCredsToken()
	if err != nil {
		t.Fatalf("Could not generate token:\n%s\n", err)
	}

	t.Logf("Token generated: %s\n", token.AccessToken)

	// Checks if token exists in the Redis cache
	res := VerifyClientCredsToken(token.AccessToken)
	if !res {
		t.Fatalf("Client Credentials token verification failed\n")
	}

	// Remove the token from the cache
	invalidateClientCredsToken(token.AccessToken)
	t.Logf("Token invalidated\n")
}

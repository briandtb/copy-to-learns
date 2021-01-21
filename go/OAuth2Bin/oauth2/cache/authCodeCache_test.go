package cache

import "testing"

// TestAuthCodeFlow tests the entirely of the functions set of authCodeStore
// as they would be used by the Authorization Code flow
func TestAuthCodeFlow(t *testing.T) {
	// Generating an authorization grant which would
	// be generated after the user authorizes the client app.
	code := NewAuthCodeGrant("https://oauth2bin.org")
	t.Logf("Generated authorization code grant: %s\n", code)

	// Generating a token based on the grant which would
	// be generated by invoking the token endpoint
	token, err := NewAuthCodeToken(code, "", "https://oauth2bin.org")
	if err != nil {
		t.Fatalf("Could not generate token:\n%s\n", err)
	}

	t.Logf("Token generated: %s\n", token.AccessToken)

	// Check if token exists
	res := VerifyAuthCodeToken(token.AccessToken)
	if !res {
		t.Fatalf("Auth Code token verification failed\n")
	}

	// Issue new token based on the previously issued refresh token
	token, err = NewAuthCodeRefreshToken(token.RefreshToken)
	if err != nil {
		t.Fatalf("Could not generate token from refresh token\n")
	}

	// Remove the token
	invalidateAuthCodeToken(token.AccessToken)
	t.Logf("Token invalidated\n")
}

func TestRefreshTokenExists(t *testing.T) {
	code := NewAuthCodeGrant("https://oauth2bin.org")
	token, err := NewAuthCodeToken(code, "", "https://oauth2bin.org")
	if err != nil {
		t.Fatal(err)
	}

	exists := AuthCodeRefreshTokenExists(token.RefreshToken, true)

	if exists {
		t.Log("found refresh token")
	} else {
		removeAuthCodeGrant(code, "https://oauth2bin.org")
		t.Fatal("failed to find refresh token")
	}
}

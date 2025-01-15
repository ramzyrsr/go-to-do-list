package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/exp/rand"
)

var (
	// secret     = os.Getenv("SECRET_SESSION")
	// sessionKey = []byte(secret)
	// sessionKey = []byte(os.Getenv("SECRET_SESSION"))
	sessionKey = []byte("super-secret-session-key")
	store      = sessions.NewCookieStore(sessionKey)
)

// SessionExpiry is the duration after which the session expires due to inactivity.
const SessionExpiry = 30 * 24 * time.Hour // 30 days
// const SessionExpiry = 5 * time.Second // 30 days

// SessionTokenExpiry defines how long the session token lasts
const SessionTokenExpiry = 24 * time.Hour // Token expiry after 24 hours of use

// GenerateSessionToken creates a new session token (UUID-like random string).
func GenerateSessionToken() string {
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tokenLength := 32
	token := make([]byte, tokenLength)
	for i := range token {
		token[i] = charset[r.Intn(len(charset))]
	}
	return string(token)
}

// CreateSession creates a new session for a user and regenerates the session ID.
func CreateSession(w http.ResponseWriter, r *http.Request, userID string) error {
	// Regenerate the session ID to prevent session fixation
	session, err := store.Get(r, "user-session")
	if err != nil {
		return fmt.Errorf("unable to create session: %v", err)
	}

	// Generate a new session token
	sessionToken := GenerateSessionToken()

	// Set session data (e.g., store user ID, or other info in the session)
	session.Values["session_token"] = sessionToken
	session.Values["user_id"] = userID
	session.Values["authenticated"] = true
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(SessionExpiry.Seconds()), // Automatically expire session after 30 days
		HttpOnly: true,                         // Make session cookie HTTP-only to prevent JS access
		Secure:   false,
	}

	// Save the session (this will update the cookie in the user's browser)
	session.Save(r, w)

	return nil
}

// DestroySession destroys the session by clearing it.
func DestroySession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "user-session")
	if err != nil {
		return fmt.Errorf("unable to get session: %v", err)
	}

	// Destroy the session by clearing it and setting MaxAge to -1
	session.Options.MaxAge = -1
	session.Save(r, w)
	return nil
}

// IsAuthenticated checks if the user is authenticated based on the session.
func IsAuthenticated(r *http.Request) (bool, error) {
	session, err := store.Get(r, "user-session")
	fmt.Println("session::", *session, "r::", r)
	if err != nil {
		return false, fmt.Errorf("unable to get session: %v", err)
	}

	// Print session values
	for key, value := range session.Values {
		fmt.Printf("Session key: %s, Value: %v\n", key, value)
	}

	// Check if session_token and authenticated values are present
	sessionToken, ok := session.Values["session_token"].(string)
	if !ok || sessionToken == "" {
		return false, nil
	}

	authenticated, ok := session.Values["authenticated"].(bool)
	if !ok || !authenticated {
		return false, nil
	}
	return true, nil
}

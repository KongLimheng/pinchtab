package testutil

import (
	"os"
	"strings"
)

// IntegrationHeaded reports whether integration tests should launch visible
// Chrome windows. Default is false to avoid desktop crash/restore dialogs
// during automated runs. Set PINCHTAB_TEST_HEADED=1 to opt in locally.
func IntegrationHeaded() bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("PINCHTAB_TEST_HEADED"))) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

package testutil

import "testing"

func TestIntegrationHeaded(t *testing.T) {
	t.Setenv("PINCHTAB_TEST_HEADED", "")
	if IntegrationHeaded() {
		t.Fatal("expected default IntegrationHeaded() to be false")
	}

	for _, value := range []string{"1", "true", "TRUE", "yes", "on"} {
		t.Setenv("PINCHTAB_TEST_HEADED", value)
		if !IntegrationHeaded() {
			t.Fatalf("expected IntegrationHeaded() true for %q", value)
		}
	}

	t.Setenv("PINCHTAB_TEST_HEADED", "0")
	if IntegrationHeaded() {
		t.Fatal("expected IntegrationHeaded() false for 0")
	}
}

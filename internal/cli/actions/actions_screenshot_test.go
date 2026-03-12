package actions

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScreenshot(t *testing.T) {
	m := newMockServer()
	m.response = "FAKEJPEGDATA"
	defer m.close()
	client := m.server.Client()

	outFile := filepath.Join(t.TempDir(), "test.jpg")
	Screenshot(client, m.base(), "", []string{"-o", outFile, "-q", "50"})
	if m.lastPath != "/screenshot" {
		t.Errorf("expected /screenshot, got %s", m.lastPath)
	}
	if !strings.Contains(m.lastQuery, "quality=50") {
		t.Errorf("expected quality=50, got %s", m.lastQuery)
	}
	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("file not written: %v", err)
	}
	if string(data) != "FAKEJPEGDATA" {
		t.Errorf("unexpected content: %s", string(data))
	}
}

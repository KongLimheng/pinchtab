package actions

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPDF(t *testing.T) {
	m := newMockServer()
	m.response = "FAKEPDFDATA"
	defer m.close()
	client := m.server.Client()

	outFile := filepath.Join(t.TempDir(), "test.pdf")
	PDF(client, m.base(), "", []string{"-o", outFile, "--tab", "tab-abc", "--landscape", "--scale", "0.8"})
	if m.lastPath != "/tabs/tab-abc/pdf" {
		t.Errorf("expected /tabs/tab-abc/pdf, got %s", m.lastPath)
	}
	if !strings.Contains(m.lastQuery, "landscape=true") {
		t.Errorf("expected landscape=true, got %s", m.lastQuery)
	}
	if !strings.Contains(m.lastQuery, "scale=0.8") {
		t.Errorf("expected scale=0.8, got %s", m.lastQuery)
	}
	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("file not written: %v", err)
	}
	if string(data) != "FAKEPDFDATA" {
		t.Errorf("unexpected content: %s", string(data))
	}
}

func TestPDFAllOptions(t *testing.T) {
	m := newMockServer()
	m.response = "FAKEPDFDATA"
	defer m.close()
	client := m.server.Client()

	outFile := filepath.Join(t.TempDir(), "test.pdf")
	args := []string{
		"-o", outFile,
		"--landscape",
		"--scale", "1.5",
		"--paper-width", "11",
		"--paper-height", "8.5",
		"--margin-top", "1",
		"--margin-bottom", "1",
		"--margin-left", "0.5",
		"--margin-right", "0.5",
		"--page-ranges", "1-3,5",
		"--prefer-css-page-size",
		"--display-header-footer",
		"--header-template", "<span class='title'></span>",
		"--footer-template", "<span class='pageNumber'></span>",
		"--generate-tagged-pdf",
		"--generate-document-outline",
		"--tab", "tab-123",
	}

	PDF(client, m.base(), "", args)
	if m.lastPath != "/tabs/tab-123/pdf" {
		t.Errorf("expected /tabs/tab-123/pdf, got %s", m.lastPath)
	}

	// Check all parameters were set correctly
	expectedParams := []string{
		"landscape=true",
		"scale=1.5",
		"paperWidth=11",
		"paperHeight=8.5",
		"marginTop=1",
		"marginBottom=1",
		"marginLeft=0.5",
		"marginRight=0.5",
		"pageRanges=1-3%2C5", // URL encoded
		"preferCSSPageSize=true",
		"displayHeaderFooter=true",
		"generateTaggedPDF=true",
		"generateDocumentOutline=true",
		"raw=true",
	}

	for _, expected := range expectedParams {
		if !strings.Contains(m.lastQuery, expected) {
			t.Errorf("expected %s in query, got %s", expected, m.lastQuery)
		}
	}
}

package actions

import (
	"strings"
	"testing"
)

func TestText(t *testing.T) {
	m := newMockServer()
	m.response = `{"url":"https://pinchtab.com","title":"Example","text":"Hello"}`
	defer m.close()
	client := m.server.Client()

	Text(client, m.base(), "", nil)
	if m.lastPath != "/text" {
		t.Errorf("expected /text, got %s", m.lastPath)
	}
}

func TestTextRaw(t *testing.T) {
	m := newMockServer()
	defer m.close()
	client := m.server.Client()

	Text(client, m.base(), "", []string{"--raw"})
	if !strings.Contains(m.lastQuery, "mode=raw") {
		t.Errorf("expected mode=raw, got %s", m.lastQuery)
	}
}

func TestTextTab(t *testing.T) {
	m := newMockServer()
	defer m.close()
	client := m.server.Client()

	Text(client, m.base(), "", []string{"--tab", "TAB1"})
	if !strings.Contains(m.lastQuery, "tabId=TAB1") {
		t.Errorf("expected tabId=TAB1, got %s", m.lastQuery)
	}
}

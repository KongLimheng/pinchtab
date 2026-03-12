package actions

import (
	"encoding/json"
	"testing"
)

func TestTabsList(t *testing.T) {
	m := newMockServer()
	m.response = `[{"id":"TAB1","url":"https://pinchtab.com"}]`
	defer m.close()
	client := m.server.Client()

	Tabs(client, m.base(), "", nil)
	if m.lastPath != "/tabs" {
		t.Errorf("expected /tabs, got %s", m.lastPath)
	}
}

func TestTabsNew(t *testing.T) {
	m := newMockServer()
	defer m.close()
	client := m.server.Client()

	Tabs(client, m.base(), "", []string{"new", "https://pinchtab.com"})
	if m.lastPath != "/tab" {
		t.Errorf("expected /tab, got %s", m.lastPath)
	}
	var body map[string]any
	_ = json.Unmarshal([]byte(m.lastBody), &body)
	if body["action"] != "new" {
		t.Errorf("expected action=new, got %v", body["action"])
	}
	if body["url"] != "https://pinchtab.com" {
		t.Errorf("expected url, got %v", body["url"])
	}
}

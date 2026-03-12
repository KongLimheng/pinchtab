package actions

import (
	"encoding/json"
	"testing"
)

func TestNavigate(t *testing.T) {
	m := newMockServer()
	defer m.close()
	client := m.server.Client()

	Navigate(client, m.base(), "", []string{"https://pinchtab.com"})
	if m.lastMethod != "POST" {
		t.Errorf("expected POST, got %s", m.lastMethod)
	}
	if m.lastPath != "/navigate" {
		t.Errorf("expected /navigate, got %s", m.lastPath)
	}
	var body map[string]any
	_ = json.Unmarshal([]byte(m.lastBody), &body)
	if body["url"] != "https://pinchtab.com" {
		t.Errorf("expected url=https://pinchtab.com, got %v", body["url"])
	}
}

func TestNavigateWithFlags(t *testing.T) {
	m := newMockServer()
	defer m.close()
	client := m.server.Client()

	Navigate(client, m.base(), "", []string{"https://pinchtab.com", "--new-tab", "--block-images"})
	var body map[string]any
	_ = json.Unmarshal([]byte(m.lastBody), &body)
	if body["newTab"] != true {
		t.Error("expected newTab=true")
	}
	if body["blockImages"] != true {
		t.Error("expected blockImages=true")
	}
}

func TestNavigateWithBlockAds(t *testing.T) {
	m := newMockServer()
	defer m.close()
	client := m.server.Client()

	Navigate(client, m.base(), "", []string{"https://pinchtab.com", "--block-ads"})
	var body map[string]any
	_ = json.Unmarshal([]byte(m.lastBody), &body)
	if body["blockAds"] != true {
		t.Error("expected blockAds=true")
	}
}

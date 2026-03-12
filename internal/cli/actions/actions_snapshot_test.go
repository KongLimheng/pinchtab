package actions

import (
	"strings"
	"testing"
)

func TestSnapshot(t *testing.T) {
	m := newMockServer()
	m.response = `[{"ref":"e0","role":"button","name":"Submit"}]`
	defer m.close()
	client := m.server.Client()

	Snapshot(client, m.base(), "", []string{"-i", "-c"})
	if m.lastMethod != "GET" {
		t.Errorf("expected GET, got %s", m.lastMethod)
	}
	if m.lastPath != "/snapshot" {
		t.Errorf("expected /snapshot, got %s", m.lastPath)
	}
	if !strings.Contains(m.lastQuery, "filter=interactive") {
		t.Errorf("expected filter=interactive in query, got %s", m.lastQuery)
	}
	if !strings.Contains(m.lastQuery, "format=compact") {
		t.Errorf("expected format=compact in query, got %s", m.lastQuery)
	}
}

func TestSnapshotDiff(t *testing.T) {
	m := newMockServer()
	defer m.close()
	client := m.server.Client()

	Snapshot(client, m.base(), "", []string{"--diff", "--selector", "main", "--max-tokens", "2000", "--depth", "5"})
	if !strings.Contains(m.lastQuery, "diff=true") {
		t.Errorf("expected diff=true, got %s", m.lastQuery)
	}
	if !strings.Contains(m.lastQuery, "selector=main") {
		t.Errorf("expected selector=main, got %s", m.lastQuery)
	}
	if !strings.Contains(m.lastQuery, "maxTokens=2000") {
		t.Errorf("expected maxTokens=2000, got %s", m.lastQuery)
	}
	if !strings.Contains(m.lastQuery, "depth=5") {
		t.Errorf("expected depth=5, got %s", m.lastQuery)
	}
}

func TestSnapshotTabId(t *testing.T) {
	m := newMockServer()
	defer m.close()
	client := m.server.Client()

	Snapshot(client, m.base(), "", []string{"--tab", "ABC123"})
	if !strings.Contains(m.lastQuery, "tabId=ABC123") {
		t.Errorf("expected tabId=ABC123, got %s", m.lastQuery)
	}
}

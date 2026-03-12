package actions

import (
	"encoding/json"
	"testing"
)

func TestEvaluate(t *testing.T) {
	m := newMockServer()
	m.response = `{"result":"Example Domain"}`
	defer m.close()
	client := m.server.Client()

	Evaluate(client, m.base(), "", []string{"document.title"})
	if m.lastPath != "/evaluate" {
		t.Errorf("expected /evaluate, got %s", m.lastPath)
	}
	var body map[string]any
	_ = json.Unmarshal([]byte(m.lastBody), &body)
	if body["expression"] != "document.title" {
		t.Errorf("expected expression=document.title, got %v", body["expression"])
	}
}

func TestEvaluateMultiWord(t *testing.T) {
	m := newMockServer()
	defer m.close()
	client := m.server.Client()

	Evaluate(client, m.base(), "", []string{"1", "+", "2"})
	var body map[string]any
	_ = json.Unmarshal([]byte(m.lastBody), &body)
	if body["expression"] != "1 + 2" {
		t.Errorf("expected expression='1 + 2', got %v", body["expression"])
	}
}

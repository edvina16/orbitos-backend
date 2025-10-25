package models

import (
	"encoding/json"
	"testing"
)

func TestTaskJSONMarshalUnmarshal(t *testing.T) {
	task := Task{
		ID:      1,
		Title:   "Test title",
		Content: "Test content",
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}

	jsonStr := string(data)
	if !(contains(jsonStr, `"id":`) && contains(jsonStr, `"title":`) && contains(jsonStr, `"content":`)) {
		t.Fatalf("Failed to marshal task: %v", jsonStr)
	}

	var unmarshaled Task
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal task: %v", err)
	}

	if unmarshaled.ID != task.ID || unmarshaled.Title != task.Title || unmarshaled.Content != task.Content {
		t.Errorf("Failed to unmarshal task: %v", unmarshaled)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) > 0 && (contains(s[1:], substr) || s[:len(substr)] == substr))
}

package core

import "testing"

func TestGetDefaultEditor(t *testing.T) {
	editor := GetDefaultEditor()
	if editor == "" {
		t.Error("No default editor detected")
	}
	t.Log("Default editor is", editor)
}

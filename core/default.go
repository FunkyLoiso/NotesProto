package core

import (
	"os"
)

const (
	envEditorStr      = "EDITOR"
	envVisualStr      = "VISUAL"
	pathEditorStr     = "/usr/bin/editor"
	pathSensEditorStr = "/usr/bin/sensible-editor"
)

// empty string if not found
func GetDefaultEditor() string {
	if val, found := os.LookupEnv(envEditorStr); found {
		return val
	}
	if val, found := os.LookupEnv(envVisualStr); found {
		return val
	}
	if _, err := os.Stat(pathEditorStr); err == nil {
		return pathEditorStr
	}
	if _, err := os.Stat(pathSensEditorStr); err == nil {
		return pathSensEditorStr
	}

	return ""
}

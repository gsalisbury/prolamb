package config

import (
	"encoding/base64"
	"testing"
)

func TestGetTargetFromRawQuery(t *testing.T) {
	json := []byte(`
{
   'type': 'http'
}`)
	val := base64.URLEncoding.EncodeToString(json)
	if err := GetTargetsFromRawQuery([]string{val}); err != nil {
		t.Fatalf("unexpected error getting targets for raw test query: %v", err)
	}
	t.Fatal("bad")
}

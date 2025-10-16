package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnderscoreToCamel(t *testing.T) {
	cases := []struct {
		Input, Expected string
	}{
		{"hello_world", "HelloWorld"},
		{"HELLO_WORLD", "HELLOWORLD"},
		{"hello__world", "HelloWorld"},
		{"helloworld", "Helloworld"},
	}

	for _, c := range cases {
		t.Run(c.Input, func(t *testing.T) {
			assert.Equal(t, c.Expected, UnderscoreToCamel(c.Input))
		})
	}
}

func TestReplaceNonASCII(t *testing.T) {
	cases := []struct {
		Input, Expected string
	}{
		{"hello_中文_world", "hello____world"},
		{"hello_world", "hello_world"},
		{"héllo_wörld", "h_llo_w_rld"},
	}

	for _, c := range cases {
		t.Run(c.Input, func(t *testing.T) {
			assert.Equal(t, c.Expected, ReplaceNonASCII(c.Input, '_'))
		})
	}
}

func TestGenerateUUIDv4(t *testing.T) {
	uuid, err := generateUUIDv4()
	assert.NoError(t, err, "Expected no error generating UUID")
	assert.Regexp(t, `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`, uuid,
		"Expected valid UUID format")

	uuid2, err := generateUUIDv4()
	assert.NoError(t, err, "Expected no error generating UUID")
	assert.NotEqual(t, uuid, uuid2, "Expected UUIDs to be unique")
}

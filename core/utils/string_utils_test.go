package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnderscoreToCamel(t *testing.T) {
	input := "hello_world"
	expected := "HelloWorld"
	result := UnderscoreToCamel(input)
	assert.Equal(t, expected, result, "Expected camel case conversion to match")

	input = "HELLO_WORLD"
	expected = "HELLOWORLD"
	result = UnderscoreToCamel(input)
	assert.Equal(t, expected, result, "Expected camel case conversion to match")

	input = "hello__world"
	expected = "HelloWorld"
	result = UnderscoreToCamel(input)
	assert.Equal(t, expected, result, "Expected camel case conversion to match")

	input = "helloworld"
	expected = "Helloworld"
	result = UnderscoreToCamel(input)
	assert.Equal(t, expected, result, "Expected camel case conversion to match")
}

func TestReplaceNonASCII(t *testing.T) {
	input := "hello_中文_world"
	replacement := '_'
	expected := "hello____world"
	result := ReplaceNonASCII(input, replacement)
	assert.Equal(t, expected, result, "Expected non-ASCII characters to be replaced")

	input = "hello_world"
	replacement = '_'
	expected = "hello_world"
	result = ReplaceNonASCII(input, replacement)
	assert.Equal(t, expected, result, "Expected no changes for ASCII-only input")

	input = "héllo_wörld"
	replacement = '_'
	expected = "h_llo_w_rld"
	result = ReplaceNonASCII(input, replacement)
	assert.Equal(t, expected, result, "Expected non-ASCII characters to be replaced")
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

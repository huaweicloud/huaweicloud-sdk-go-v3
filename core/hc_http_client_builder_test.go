package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHcHttpClientBuilder_WithEndpoint(t *testing.T) {
	builder := NewHcHttpClientBuilder().WithEndpoint("endpoint")
	assert.Equal(t, []string{"endpoint"}, builder.endpoints)

	builder = NewHcHttpClientBuilder().WithEndpoint("endpoint1", "endpoint2")
	assert.Equal(t, []string{"endpoint1", "endpoint2"}, builder.endpoints)
}

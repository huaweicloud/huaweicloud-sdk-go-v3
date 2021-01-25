package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceShowImageSchemasRequest struct {
}

func (o GlanceShowImageSchemasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceShowImageSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageSchemasRequest", string(data)}, " ")
}

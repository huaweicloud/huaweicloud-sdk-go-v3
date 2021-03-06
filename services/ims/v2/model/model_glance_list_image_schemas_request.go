package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceListImageSchemasRequest struct {
}

func (o GlanceListImageSchemasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceListImageSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageSchemasRequest", string(data)}, " ")
}

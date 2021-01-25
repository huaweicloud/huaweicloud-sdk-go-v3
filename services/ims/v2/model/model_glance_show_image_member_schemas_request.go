package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceShowImageMemberSchemasRequest struct {
}

func (o GlanceShowImageMemberSchemasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberSchemasRequest", string(data)}, " ")
}

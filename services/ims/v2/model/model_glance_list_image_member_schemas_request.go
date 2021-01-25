package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceListImageMemberSchemasRequest struct {
}

func (o GlanceListImageMemberSchemasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceListImageMemberSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageMemberSchemasRequest", string(data)}, " ")
}

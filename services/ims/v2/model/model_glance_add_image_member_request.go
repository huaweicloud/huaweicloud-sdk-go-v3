package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceAddImageMemberRequest struct {
	ImageId string                           `json:"image_id"`
	Body    *GlanceAddImageMemberRequestBody `json:"body,omitempty"`
}

func (o GlanceAddImageMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceAddImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceAddImageMemberRequest", string(data)}, " ")
}

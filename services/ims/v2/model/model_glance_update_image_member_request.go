package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceUpdateImageMemberRequest struct {
	ImageId string `json:"image_id"`

	MemberId string `json:"member_id"`

	Body *GlanceUpdateImageMemberRequestBody `json:"body,omitempty"`
}

func (o GlanceUpdateImageMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageMemberRequest", string(data)}, " ")
}

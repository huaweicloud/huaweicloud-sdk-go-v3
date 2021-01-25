package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceDeleteImageMemberRequest struct {
	ImageId  string `json:"image_id"`
	MemberId string `json:"member_id"`
}

func (o GlanceDeleteImageMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageMemberRequest", string(data)}, " ")
}

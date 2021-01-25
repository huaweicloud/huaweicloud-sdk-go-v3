package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceListImageMembersRequest struct {
	ImageId string `json:"image_id"`
}

func (o GlanceListImageMembersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceListImageMembersRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageMembersRequest", string(data)}, " ")
}

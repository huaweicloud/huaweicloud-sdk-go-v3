package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceListImageMembersRequest struct {
	// 镜像id

	ImageId string `json:"image_id"`
}

func (o GlanceListImageMembersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceListImageMembersRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageMembersRequest", string(data)}, " ")
}

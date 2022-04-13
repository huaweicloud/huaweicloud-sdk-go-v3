package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateGrantRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *CreateGrantRequestBody `json:"body,omitempty"`
}

func (o CreateGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGrantRequest struct{}"
	}

	return strings.Join([]string{"CreateGrantRequest", string(data)}, " ")
}

package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDatakeyRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *CreateDatakeyRequestBody `json:"body,omitempty"`
}

func (o CreateDatakeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyRequest struct{}"
	}

	return strings.Join([]string{"CreateDatakeyRequest", string(data)}, " ")
}

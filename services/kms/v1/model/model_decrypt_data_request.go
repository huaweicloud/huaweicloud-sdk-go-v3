package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DecryptDataRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *DecryptDataRequestBody `json:"body,omitempty"`
}

func (o DecryptDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDataRequest struct{}"
	}

	return strings.Join([]string{"DecryptDataRequest", string(data)}, " ")
}

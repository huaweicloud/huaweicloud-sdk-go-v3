package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListKeysRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *ListKeysRequestBody `json:"body,omitempty"`
}

func (o ListKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeysRequest struct{}"
	}

	return strings.Join([]string{"ListKeysRequest", string(data)}, " ")
}

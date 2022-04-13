package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListGrantsRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *ListGrantsRequestBody `json:"body,omitempty"`
}

func (o ListGrantsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGrantsRequest struct{}"
	}

	return strings.Join([]string{"ListGrantsRequest", string(data)}, " ")
}

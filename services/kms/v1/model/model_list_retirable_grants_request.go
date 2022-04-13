package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRetirableGrantsRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *ListRetirableGrantsRequestBody `json:"body,omitempty"`
}

func (o ListRetirableGrantsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetirableGrantsRequest struct{}"
	}

	return strings.Join([]string{"ListRetirableGrantsRequest", string(data)}, " ")
}

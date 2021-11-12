package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateBaremetalServerTagsRequest struct {
	// 裸金属服务器ID。

	ServerId string `json:"server_id"`

	Body *BatchCreateBaremetalServerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateBaremetalServerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateBaremetalServerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateBaremetalServerTagsRequest", string(data)}, " ")
}

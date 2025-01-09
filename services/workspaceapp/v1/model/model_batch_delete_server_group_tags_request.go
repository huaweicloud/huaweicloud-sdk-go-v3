package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServerGroupTagsRequest Request Object
type BatchDeleteServerGroupTagsRequest struct {
	Body *BatchDeleteServerGroupTagsReq `json:"body,omitempty"`
}

func (o BatchDeleteServerGroupTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerGroupTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerGroupTagsRequest", string(data)}, " ")
}

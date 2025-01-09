package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateServerGroupTagsRequest Request Object
type BatchCreateServerGroupTagsRequest struct {
	Body *BatchCreateServerGroupTagsReq `json:"body,omitempty"`
}

func (o BatchCreateServerGroupTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServerGroupTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateServerGroupTagsRequest", string(data)}, " ")
}

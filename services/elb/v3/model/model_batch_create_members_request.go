package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateMembersRequest struct {
	// 后端服务器组ID。

	PoolId string `json:"pool_id"`

	Body *BatchCreateMembersRequestBody `json:"body,omitempty"`
}

func (o BatchCreateMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersRequest", string(data)}, " ")
}

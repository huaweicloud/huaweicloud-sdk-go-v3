package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteMembersRequest struct {

	// 后端服务器组ID。
	PoolId string `json:"pool_id" xml:"pool_id"`

	Body *BatchDeleteMembersRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchDeleteMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersRequest", string(data)}, " ")
}

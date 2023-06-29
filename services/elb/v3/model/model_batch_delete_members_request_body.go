package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteMembersRequestBody This is a auto create Body Object
type BatchDeleteMembersRequestBody struct {

	// 批量删除后端服务器请求body。
	Members []BatchDeleteMembersOption `json:"members"`
}

func (o BatchDeleteMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersRequestBody", string(data)}, " ")
}

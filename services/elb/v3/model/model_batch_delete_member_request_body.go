package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type BatchDeleteMemberRequestBody struct {
	// 批量删除后端服务器请求body。

	Members []BatchDeleteMembersOption `json:"members"`
}

func (o BatchDeleteMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMemberRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteMemberRequestBody", string(data)}, " ")
}

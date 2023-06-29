package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteMemberReq 批量删除用户请求体
type BatchDeleteMemberReq struct {

	// 删除用户列表
	Members []MemberDto `json:"members"`
}

func (o BatchDeleteMemberReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMemberReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteMemberReq", string(data)}, " ")
}

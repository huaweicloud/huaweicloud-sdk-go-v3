package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteMemberRsp 批量删除用户返回体
type BatchDeleteMemberRsp struct {

	// 用户id
	Id *string `json:"id,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`

	// 删除结果
	Status *string `json:"status,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o BatchDeleteMemberRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMemberRsp struct{}"
	}

	return strings.Join([]string{"BatchDeleteMemberRsp", string(data)}, " ")
}

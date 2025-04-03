package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetGroupResp struct {

	// 用户组的全局唯一标识符（ID）
	Id string `json:"id"`

	// 外部标识符
	ExternalId *string `json:"externalId,omitempty"`

	Meta *MetaDto `json:"meta,omitempty"`

	// 概要
	Schemas *[]string `json:"schemas,omitempty"`

	// 包含用户显示名称的字符串
	DisplayName *string `json:"displayName,omitempty"`

	// 用户组中的成员对象列表
	Members *[]MemberItemDto `json:"members,omitempty"`
}

func (o GetGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetGroupResp struct{}"
	}

	return strings.Join([]string{"GetGroupResp", string(data)}, " ")
}

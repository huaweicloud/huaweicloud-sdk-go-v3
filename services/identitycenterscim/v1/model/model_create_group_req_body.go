package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGroupReqBody struct {

	// 外部标识符
	ExternalId *string `json:"externalId,omitempty"`

	// 包含用户组显示名称的字符串
	DisplayName string `json:"displayName"`

	// 用户组中的成员对象列表
	Members *[]MemberItemDto `json:"members,omitempty"`

	// 概要
	Schemas []string `json:"schemas"`
}

func (o CreateGroupReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupReqBody struct{}"
	}

	return strings.Join([]string{"CreateGroupReqBody", string(data)}, " ")
}

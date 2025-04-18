package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupResponse Response Object
type CreateGroupResponse struct {

	// 用户组的全局唯一标识符（ID）
	Id *string `json:"id,omitempty"`

	// 外部标识符
	ExternalId *string `json:"externalId,omitempty"`

	Meta *MetaDto `json:"meta,omitempty"`

	// 概要
	Schemas *[]string `json:"schemas,omitempty"`

	// 包含用户组显示名称的字符串
	DisplayName *string `json:"displayName,omitempty"`

	// 用户组中的成员对象列表
	Members        *[]MemberItemDto `json:"members,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateGroupResponse", string(data)}, " ")
}

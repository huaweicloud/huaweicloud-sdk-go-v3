package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberItemDto struct {

	// 成员的全局唯一标识符（ID）
	Value string `json:"value"`

	// 成员的引用信息
	Ref *string `json:"$ref,omitempty"`

	// 成员类型 User：用户
	Type *string `json:"type,omitempty"`
}

func (o MemberItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberItemDto struct{}"
	}

	return strings.Join([]string{"MemberItemDto", string(data)}, " ")
}

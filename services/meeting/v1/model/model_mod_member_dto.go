package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModMemberDto 用户自己修改个人信息。
type ModMemberDto struct {

	// 名称。
	Name string `json:"name"`

	// 英文名称。
	EnglishName *string `json:"englishName,omitempty"`

	// 签名。
	Signature *string `json:"signature,omitempty"`

	// 职位。
	Title *string `json:"title,omitempty"`

	// 备注。
	Desc *string `json:"desc,omitempty"`
}

func (o ModMemberDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModMemberDto struct{}"
	}

	return strings.Join([]string{"ModMemberDto", string(data)}, " ")
}

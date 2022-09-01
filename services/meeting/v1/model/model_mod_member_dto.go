package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户自己修改个人信息的DTO，用member命名做区分，当前仅有名称，后续会扩展地址、签名等等。
type ModMemberDto struct {

	// 名称 maxLength：64 minLength：1
	Name string `json:"name" xml:"name"`

	// 英文名称 maxLength：64 minLength：0
	EnglishName *string `json:"englishName,omitempty" xml:"englishName"`

	// 签名 maxLength：512 minLength：0
	Signature *string `json:"signature,omitempty" xml:"signature"`

	// 职位 maxLength：32 minLength：0
	Title *string `json:"title,omitempty" xml:"title"`

	// 备注 maxLength：128 minLength：0
	Desc *string `json:"desc,omitempty" xml:"desc"`
}

func (o ModMemberDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModMemberDto struct{}"
	}

	return strings.Join([]string{"ModMemberDto", string(data)}, " ")
}

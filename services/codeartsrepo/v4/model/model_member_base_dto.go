package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberBaseDto 成员基础信息
type MemberBaseDto struct {

	// 用户id
	Id *string `json:"id,omitempty"`

	// iam_id
	IamId *string `json:"iam_id,omitempty"`

	// 成员名字
	Name *string `json:"name,omitempty"`
}

func (o MemberBaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberBaseDto struct{}"
	}

	return strings.Join([]string{"MemberBaseDto", string(data)}, " ")
}

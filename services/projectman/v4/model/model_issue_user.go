package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueUser struct {

	// 用户id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 用户名
	Name *string `json:"name,omitempty" xml:"name"`

	// 昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`
}

func (o IssueUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueUser struct{}"
	}

	return strings.Join([]string{"IssueUser", string(data)}, " ")
}

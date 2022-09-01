package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作项操作的用户
type IssueRecordV4User struct {

	// 用户数字id
	UserNumId *int32 `json:"user_num_id,omitempty" xml:"user_num_id"`

	// 登录名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`

	// 用户32位的uuid
	UserId *string `json:"user_id,omitempty" xml:"user_id"`
}

func (o IssueRecordV4User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueRecordV4User struct{}"
	}

	return strings.Join([]string{"IssueRecordV4User", string(data)}, " ")
}

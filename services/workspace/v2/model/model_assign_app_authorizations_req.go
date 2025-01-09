package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssignAppAuthorizationsReq 设置应用授权。
type AssignAppAuthorizationsReq struct {
	AuthorizationType *AssignType `json:"authorization_type"`

	// 新增授权用户列表，一次请求数量区间 [0, 100]。
	Users *[]AccountInfo `json:"users,omitempty"`

	// 取消授权用户列表，一次请求数量区间 [0, 100]。
	DelUsers *[]AccountInfo `json:"del_users,omitempty"`
}

func (o AssignAppAuthorizationsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssignAppAuthorizationsReq struct{}"
	}

	return strings.Join([]string{"AssignAppAuthorizationsReq", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssignAppAuthorizationsReq 批量设置应用授权。
type BatchAssignAppAuthorizationsReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 20]。
	AppIds []string `json:"app_ids"`

	AuthorizationType *AssignType `json:"authorization_type"`

	// 新增授权用户列表，一次请求数量区间 [0, 100]。
	Users *[]AccountInfo `json:"users,omitempty"`

	// 取消授权用户列表，一次请求数量区间 [0, 100]。
	DelUsers *[]AccountInfo `json:"del_users,omitempty"`
}

func (o BatchAssignAppAuthorizationsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssignAppAuthorizationsReq struct{}"
	}

	return strings.Join([]string{"BatchAssignAppAuthorizationsReq", string(data)}, " ")
}

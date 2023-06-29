package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSelfPrivilegesResponse Response Object
type ListSelfPrivilegesResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 权限列表信息
	UserPrivileges *[]UserPrivilegeInfo `json:"user_privileges,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListSelfPrivilegesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSelfPrivilegesResponse struct{}"
	}

	return strings.Join([]string{"ListSelfPrivilegesResponse", string(data)}, " ")
}

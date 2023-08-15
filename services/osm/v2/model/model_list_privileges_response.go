package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivilegesResponse Response Object
type ListPrivilegesResponse struct {

	// 是否有权限
	HasPrivilege   *int32 `json:"has_privilege,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPrivilegesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivilegesResponse struct{}"
	}

	return strings.Join([]string{"ListPrivilegesResponse", string(data)}, " ")
}

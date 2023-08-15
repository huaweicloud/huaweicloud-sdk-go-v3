package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserRoleReq 更新用户请求体
type UpdateUserRoleReq struct {
	Role *UserRole `json:"role"`
}

func (o UpdateUserRoleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserRoleReq struct{}"
	}

	return strings.Join([]string{"UpdateUserRoleReq", string(data)}, " ")
}

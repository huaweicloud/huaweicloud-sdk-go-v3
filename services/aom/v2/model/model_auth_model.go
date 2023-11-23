package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthModel struct {

	// 状态
	Status *bool `json:"status,omitempty"`

	// 角色列表
	RoleName *[]string `json:"role_name,omitempty"`
}

func (o AuthModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthModel struct{}"
	}

	return strings.Join([]string{"AuthModel", string(data)}, " ")
}

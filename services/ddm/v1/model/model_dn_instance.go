package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DnInstance struct {

	// 实例id。
	DnInstanceId *string `json:"dn_instance_id,omitempty"`

	// 实例账号。
	AdminUser *string `json:"admin_user,omitempty"`

	// 实例密码。
	AdminPassword *string `json:"admin_password,omitempty"`
}

func (o DnInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnInstance struct{}"
	}

	return strings.Join([]string{"DnInstance", string(data)}, " ")
}

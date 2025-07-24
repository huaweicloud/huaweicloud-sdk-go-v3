package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstancePasswordOptions 修改密码请求参数
type ChangeInstancePasswordOptions struct {

	// 设置实例的管理员账户初始登录密码，其中，Linux管理员账户为root，Windows管理员账户为Administrator。
	NewPassword string `json:"new_password"`

	// instance id set
	InstanceIdSet *[]string `json:"instance_id_set,omitempty"`
}

func (o ChangeInstancePasswordOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstancePasswordOptions struct{}"
	}

	return strings.Join([]string{"ChangeInstancePasswordOptions", string(data)}, " ")
}

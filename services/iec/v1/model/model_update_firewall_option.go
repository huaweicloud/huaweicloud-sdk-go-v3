package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新网络ACL参数。
type UpdateFirewallOption struct {

	// 网络ACL名称。更新时name不能为空。  中文字符、字母、数字、中划线和下划线组成，长度为1~64个字符
	Name string `json:"name" xml:"name"`

	// 网络ACL的使能开关。  取值范围：true（开启），false（关闭）
	AdminStateUp *bool `json:"admin_state_up,omitempty" xml:"admin_state_up"`

	// 网络ACL描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 关联子网列表。
	Subnets *[]FirewallSubnetOption `json:"subnets,omitempty" xml:"subnets"`
}

func (o UpdateFirewallOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallOption struct{}"
	}

	return strings.Join([]string{"UpdateFirewallOption", string(data)}, " ")
}

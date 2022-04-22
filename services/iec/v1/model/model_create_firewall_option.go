package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建防火墙参数
type CreateFirewallOption struct {

	// 中文字符、字母、数字、中划线和下划线组成，长度为1~64个字符
	Name string `json:"name"`

	// 网络ACL描述。  取值范围：0-64
	Description *string `json:"description,omitempty"`
}

func (o CreateFirewallOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallOption struct{}"
	}

	return strings.Join([]string{"CreateFirewallOption", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WafPolicyBindHost 绑定的域名信息
type WafPolicyBindHost struct {

	// 域名ID
	Id *string `json:"id,omitempty"`

	// 域名
	Hostname *string `json:"hostname,omitempty"`
}

func (o WafPolicyBindHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafPolicyBindHost struct{}"
	}

	return strings.Join([]string{"WafPolicyBindHost", string(data)}, " ")
}

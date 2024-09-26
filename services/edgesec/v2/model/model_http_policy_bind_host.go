package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpPolicyBindHost 绑定的域名信息
type HttpPolicyBindHost struct {

	// 域名ID
	Id *string `json:"id,omitempty"`

	// 域名
	Hostname *string `json:"hostname,omitempty"`
}

func (o HttpPolicyBindHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpPolicyBindHost struct{}"
	}

	return strings.Join([]string{"HttpPolicyBindHost", string(data)}, " ")
}

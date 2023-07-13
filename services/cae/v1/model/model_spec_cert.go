package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpecCert struct {

	// 证书内容。
	Crt *string `json:"crt,omitempty"`

	// 私钥内容。
	Key *string `json:"key,omitempty"`

	// 安全策略。 - tls-1-2-strict - tls-1-2 - tls-1-1 - tls-1-0
	Policy *string `json:"policy,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o SpecCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecCert struct{}"
	}

	return strings.Join([]string{"SpecCert", string(data)}, " ")
}

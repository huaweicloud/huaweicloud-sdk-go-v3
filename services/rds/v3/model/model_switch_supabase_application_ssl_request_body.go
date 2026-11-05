package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSupabaseApplicationSslRequestBody SSL开关请求体。
type SwitchSupabaseApplicationSslRequestBody struct {

	// 是否开启SSL。false时cert_id必须为空；true时cert_id必填。
	EnableSsl bool `json:"enable_ssl"`

	// 证书ID。enable_ssl=true时必填。
	CertId *string `json:"cert_id,omitempty"`
}

func (o SwitchSupabaseApplicationSslRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSupabaseApplicationSslRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchSupabaseApplicationSslRequestBody", string(data)}, " ")
}

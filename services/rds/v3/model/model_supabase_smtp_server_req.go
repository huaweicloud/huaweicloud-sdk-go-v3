package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupabaseSmtpServerReq struct {
	SmtpHost *string `json:"smtp_host,omitempty"`

	SmtpPort *int32 `json:"smtp_port,omitempty"`

	SmtpUser *string `json:"smtp_user,omitempty"`

	// 秘密，受控输入，show不返回。
	SmtpPwd *string `json:"smtp_pwd,omitempty"`

	SmtpAdminEmail *string `json:"smtp_admin_email,omitempty"`

	SmtpSenderName *string `json:"smtp_sender_name,omitempty"`

	SmtpMaxFrequency *int32 `json:"smtp_max_frequency,omitempty"`
}

func (o SupabaseSmtpServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseSmtpServerReq struct{}"
	}

	return strings.Join([]string{"SupabaseSmtpServerReq", string(data)}, " ")
}

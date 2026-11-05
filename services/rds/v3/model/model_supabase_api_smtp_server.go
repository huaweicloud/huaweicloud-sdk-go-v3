package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseApiSmtpServer 查询响应不含smtp_pwd。
type SupabaseApiSmtpServer struct {
	SmtpHost *string `json:"smtp_host,omitempty"`

	SmtpPort *int32 `json:"smtp_port,omitempty"`

	SmtpUser *string `json:"smtp_user,omitempty"`

	SmtpAdminEmail *string `json:"smtp_admin_email,omitempty"`

	SmtpSenderName *string `json:"smtp_sender_name,omitempty"`

	SmtpMaxFrequency *int32 `json:"smtp_max_frequency,omitempty"`
}

func (o SupabaseApiSmtpServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseApiSmtpServer struct{}"
	}

	return strings.Join([]string{"SupabaseApiSmtpServer", string(data)}, " ")
}

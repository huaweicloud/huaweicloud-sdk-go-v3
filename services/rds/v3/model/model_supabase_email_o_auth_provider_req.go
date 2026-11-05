package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseEmailOAuthProviderReq enable_email_sign_up=false时清空enable_email_auto_confirm和smtp_config。
type SupabaseEmailOAuthProviderReq struct {
	EnableEmailSignUp *bool `json:"enable_email_sign_up,omitempty"`

	EnableEmailAutoConfirm *bool `json:"enable_email_auto_confirm,omitempty"`

	SmtpConfig *SupabaseSmtpServerReq `json:"smtp_config,omitempty"`
}

func (o SupabaseEmailOAuthProviderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseEmailOAuthProviderReq struct{}"
	}

	return strings.Join([]string{"SupabaseEmailOAuthProviderReq", string(data)}, " ")
}

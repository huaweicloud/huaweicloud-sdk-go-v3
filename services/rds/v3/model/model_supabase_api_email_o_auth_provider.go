package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupabaseApiEmailOAuthProvider struct {
	EnableEmailSignUp *bool `json:"enable_email_sign_up,omitempty"`

	EnableEmailAutoConfirm *bool `json:"enable_email_auto_confirm,omitempty"`

	SmtpConfig *SupabaseApiSmtpServer `json:"smtp_config,omitempty"`
}

func (o SupabaseApiEmailOAuthProvider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseApiEmailOAuthProvider struct{}"
	}

	return strings.Join([]string{"SupabaseApiEmailOAuthProvider", string(data)}, " ")
}

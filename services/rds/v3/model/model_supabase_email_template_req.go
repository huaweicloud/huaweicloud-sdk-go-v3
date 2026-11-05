package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupabaseEmailTemplateReq struct {
	MailerConfirmation *SupabaseMailerItemTemplateReq `json:"mailer_confirmation,omitempty"`
}

func (o SupabaseEmailTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseEmailTemplateReq struct{}"
	}

	return strings.Join([]string{"SupabaseEmailTemplateReq", string(data)}, " ")
}

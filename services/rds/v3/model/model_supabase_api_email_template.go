package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupabaseApiEmailTemplate struct {
	MailerConfirmation *SupabaseMailerItemTemplateReq `json:"mailer_confirmation,omitempty"`
}

func (o SupabaseApiEmailTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseApiEmailTemplate struct{}"
	}

	return strings.Join([]string{"SupabaseApiEmailTemplate", string(data)}, " ")
}

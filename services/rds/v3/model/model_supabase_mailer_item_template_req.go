package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupabaseMailerItemTemplateReq struct {
	Subject *string `json:"subject,omitempty"`

	Url *string `json:"url,omitempty"`
}

func (o SupabaseMailerItemTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseMailerItemTemplateReq struct{}"
	}

	return strings.Join([]string{"SupabaseMailerItemTemplateReq", string(data)}, " ")
}

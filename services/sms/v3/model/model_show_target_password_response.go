package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTargetPasswordResponse Response Object
type ShowTargetPasswordResponse struct {

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 目的端密码
	TargetPassword *string `json:"target_password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTargetPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTargetPasswordResponse struct{}"
	}

	return strings.Join([]string{"ShowTargetPasswordResponse", string(data)}, " ")
}

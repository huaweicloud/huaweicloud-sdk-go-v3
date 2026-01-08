package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyTemplateResponse Response Object
type CreatePolicyTemplateResponse struct {

	// 策略组ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePolicyTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreatePolicyTemplateResponse", string(data)}, " ")
}

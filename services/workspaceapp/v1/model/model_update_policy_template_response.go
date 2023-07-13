package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyTemplateResponse Response Object
type UpdatePolicyTemplateResponse struct {

	// 被修改策略模板主键
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePolicyTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyTemplateResponse", string(data)}, " ")
}

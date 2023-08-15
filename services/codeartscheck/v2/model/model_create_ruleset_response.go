package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRulesetResponse Response Object
type CreateRulesetResponse struct {

	// 规则集ID
	TemplateId     *string `json:"template_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRulesetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRulesetResponse struct{}"
	}

	return strings.Join([]string{"CreateRulesetResponse", string(data)}, " ")
}

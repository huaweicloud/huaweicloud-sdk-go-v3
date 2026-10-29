package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRulesetResponse Response Object
type CreateRulesetResponse struct {

	// **参数解释:** 规则集ID。 **取值范围:** 32位UUID。
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

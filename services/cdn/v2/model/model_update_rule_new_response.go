package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleNewResponse Response Object
type UpdateRuleNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRuleNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleNewResponse struct{}"
	}

	return strings.Join([]string{"UpdateRuleNewResponse", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteRulesetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRulesetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRulesetResponse struct{}"
	}

	return strings.Join([]string{"DeleteRulesetResponse", string(data)}, " ")
}

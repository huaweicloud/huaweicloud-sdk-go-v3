package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteRoutingRuleResponse struct {
	Body           *string `json:"body,omitempty" xml:"body"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRoutingRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRoutingRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRoutingRuleResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWafWhiteIpRuleResponse Response Object
type ListWafWhiteIpRuleResponse struct {

	// total
	Total *int32 `json:"total,omitempty"`

	// black
	Black *[]BlackWhiteListRule `json:"black,omitempty"`

	// white
	White          *[]BlackWhiteListRule `json:"white,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListWafWhiteIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafWhiteIpRuleResponse struct{}"
	}

	return strings.Join([]string{"ListWafWhiteIpRuleResponse", string(data)}, " ")
}

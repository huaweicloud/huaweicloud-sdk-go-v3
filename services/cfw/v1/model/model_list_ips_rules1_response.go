package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsRules1Response Response Object
type ListIpsRules1Response struct {
	Data           *IpsRuleListVo `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListIpsRules1Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsRules1Response struct{}"
	}

	return strings.Join([]string{"ListIpsRules1Response", string(data)}, " ")
}

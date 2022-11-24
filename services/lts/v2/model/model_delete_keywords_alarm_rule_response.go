package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteKeywordsAlarmRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteKeywordsAlarmRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeywordsAlarmRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteKeywordsAlarmRuleResponse", string(data)}, " ")
}

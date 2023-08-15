package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmRuleResourcesResponse Response Object
type DeleteAlarmRuleResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlarmRuleResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRuleResourcesResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRuleResourcesResponse", string(data)}, " ")
}

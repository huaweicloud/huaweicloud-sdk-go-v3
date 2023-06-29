package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAlarmRuleResourcesResponse Response Object
type AddAlarmRuleResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddAlarmRuleResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAlarmRuleResourcesResponse struct{}"
	}

	return strings.Join([]string{"AddAlarmRuleResourcesResponse", string(data)}, " ")
}

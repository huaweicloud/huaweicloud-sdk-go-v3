package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAssociationAlarmsResponseAlarms struct {

	// 告警规则ID
	AlarmId string `json:"alarm_id"`

	// 告警规则名称
	Name string `json:"name"`

	// 告警规则描述
	Description string `json:"description"`
}

func (o ListAssociationAlarmsResponseAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociationAlarmsResponseAlarms struct{}"
	}

	return strings.Join([]string{"ListAssociationAlarmsResponseAlarms", string(data)}, " ")
}

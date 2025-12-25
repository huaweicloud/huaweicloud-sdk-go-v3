package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapAlertName 告警报名称
type IsapAlertName struct {
}

func (o IsapAlertName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapAlertName struct{}"
	}

	return strings.Join([]string{"IsapAlertName", string(data)}, " ")
}

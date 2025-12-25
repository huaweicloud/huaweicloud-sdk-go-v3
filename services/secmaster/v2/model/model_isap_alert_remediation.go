package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapAlertRemediation 告警修复建议
type IsapAlertRemediation struct {
}

func (o IsapAlertRemediation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapAlertRemediation struct{}"
	}

	return strings.Join([]string{"IsapAlertRemediation", string(data)}, " ")
}

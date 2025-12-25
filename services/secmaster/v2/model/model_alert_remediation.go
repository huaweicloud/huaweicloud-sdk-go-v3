package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertRemediation 告警修复
type AlertRemediation struct {
}

func (o AlertRemediation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRemediation struct{}"
	}

	return strings.Join([]string{"AlertRemediation", string(data)}, " ")
}

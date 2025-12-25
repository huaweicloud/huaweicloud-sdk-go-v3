package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertSuppression 告警抑制
type AlertSuppression struct {
}

func (o AlertSuppression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertSuppression struct{}"
	}

	return strings.Join([]string{"AlertSuppression", string(data)}, " ")
}

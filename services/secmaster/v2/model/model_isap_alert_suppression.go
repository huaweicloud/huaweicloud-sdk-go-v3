package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapAlertSuppression 抑制标志
type IsapAlertSuppression struct {
}

func (o IsapAlertSuppression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapAlertSuppression struct{}"
	}

	return strings.Join([]string{"IsapAlertSuppression", string(data)}, " ")
}

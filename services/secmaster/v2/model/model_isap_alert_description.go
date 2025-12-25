package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapAlertDescription 告警描述
type IsapAlertDescription struct {
}

func (o IsapAlertDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapAlertDescription struct{}"
	}

	return strings.Join([]string{"IsapAlertDescription", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapAlertGrouping 分组标志
type IsapAlertGrouping struct {
}

func (o IsapAlertGrouping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapAlertGrouping struct{}"
	}

	return strings.Join([]string{"IsapAlertGrouping", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertGrouping 告警组
type AlertGrouping struct {
}

func (o AlertGrouping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertGrouping struct{}"
	}

	return strings.Join([]string{"AlertGrouping", string(data)}, " ")
}

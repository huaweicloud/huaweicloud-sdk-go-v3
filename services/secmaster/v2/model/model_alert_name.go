package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertName 告警名称
type AlertName struct {
}

func (o AlertName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertName struct{}"
	}

	return strings.Join([]string{"AlertName", string(data)}, " ")
}

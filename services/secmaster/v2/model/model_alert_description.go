package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertDescription 告警描述
type AlertDescription struct {
}

func (o AlertDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertDescription struct{}"
	}

	return strings.Join([]string{"AlertDescription", string(data)}, " ")
}

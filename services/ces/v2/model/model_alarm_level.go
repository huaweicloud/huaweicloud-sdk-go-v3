package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmLevel 告警级别，1为紧急，2为重要，3为次要，4为提示
type AlarmLevel struct {
}

func (o AlarmLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmLevel struct{}"
	}

	return strings.Join([]string{"AlarmLevel", string(data)}, " ")
}

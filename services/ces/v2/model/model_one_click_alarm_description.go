package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneClickAlarmDescription 一键告警描述，长度范围[0,256]，该字段默认值为空字符串
type OneClickAlarmDescription struct {
}

func (o OneClickAlarmDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneClickAlarmDescription struct{}"
	}

	return strings.Join([]string{"OneClickAlarmDescription", string(data)}, " ")
}

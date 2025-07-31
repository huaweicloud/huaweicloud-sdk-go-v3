package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneClickAlarmEnabled 是否启用一键告警。true:开启，false：关闭。
type OneClickAlarmEnabled struct {
}

func (o OneClickAlarmEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneClickAlarmEnabled struct{}"
	}

	return strings.Join([]string{"OneClickAlarmEnabled", string(data)}, " ")
}

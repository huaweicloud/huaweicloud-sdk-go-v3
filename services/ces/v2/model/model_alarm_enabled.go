package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmEnabled 是否开启告警规则。true:开启，false:关闭。
type AlarmEnabled struct {
}

func (o AlarmEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmEnabled struct{}"
	}

	return strings.Join([]string{"AlarmEnabled", string(data)}, " ")
}

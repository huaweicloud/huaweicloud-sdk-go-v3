package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmEnabled 告警开关
type AlarmEnabled struct {
}

func (o AlarmEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmEnabled struct{}"
	}

	return strings.Join([]string{"AlarmEnabled", string(data)}, " ")
}

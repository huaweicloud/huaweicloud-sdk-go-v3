package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警规则id，以al开头，包含22个数字或字母
type AlarmId struct {
}

func (o AlarmId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmId struct{}"
	}

	return strings.Join([]string{"AlarmId", string(data)}, " ")
}

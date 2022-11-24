package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
type AlarmName struct {
}

func (o AlarmName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmName struct{}"
	}

	return strings.Join([]string{"AlarmName", string(data)}, " ")
}

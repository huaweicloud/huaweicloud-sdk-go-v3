package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警描述，长度0-256
type AlarmDescription struct {
}

func (o AlarmDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmDescription struct{}"
	}

	return strings.Join([]string{"AlarmDescription", string(data)}, " ")
}

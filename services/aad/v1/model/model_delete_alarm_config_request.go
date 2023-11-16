package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmConfigRequest Request Object
type DeleteAlarmConfigRequest struct {
}

func (o DeleteAlarmConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmConfigRequest", string(data)}, " ")
}

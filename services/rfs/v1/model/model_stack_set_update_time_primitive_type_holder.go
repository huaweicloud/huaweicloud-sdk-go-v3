package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetUpdateTimePrimitiveTypeHolder struct {

	// 资源栈集的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o StackSetUpdateTimePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetUpdateTimePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetUpdateTimePrimitiveTypeHolder", string(data)}, " ")
}

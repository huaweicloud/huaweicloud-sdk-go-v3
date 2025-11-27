package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetOperationUpdateTimePrimitiveTypeHolder struct {

	// 资源栈集操作的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o StackSetOperationUpdateTimePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetOperationUpdateTimePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetOperationUpdateTimePrimitiveTypeHolder", string(data)}, " ")
}

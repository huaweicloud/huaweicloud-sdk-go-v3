package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateHookUpdateTimePrimitiveTypeHolder struct {

	// 私有Hook（private-hook）的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o PrivateHookUpdateTimePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateHookUpdateTimePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateHookUpdateTimePrimitiveTypeHolder", string(data)}, " ")
}

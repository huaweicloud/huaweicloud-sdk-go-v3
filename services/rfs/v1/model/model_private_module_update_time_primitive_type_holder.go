package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateModuleUpdateTimePrimitiveTypeHolder struct {

	// 私有模块（private-module）的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o PrivateModuleUpdateTimePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateModuleUpdateTimePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateModuleUpdateTimePrimitiveTypeHolder", string(data)}, " ")
}

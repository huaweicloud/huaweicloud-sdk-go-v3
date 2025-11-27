package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateHookCreateTimePrimitiveTypeHolder struct {

	// 私有Hook（private-hook）的生成时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o PrivateHookCreateTimePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateHookCreateTimePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateHookCreateTimePrimitiveTypeHolder", string(data)}, " ")
}

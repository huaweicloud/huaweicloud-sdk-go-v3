package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateProviderUpdateTimePrimitiveTypeHolder struct {

	// 私有provider（private-provider）的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o PrivateProviderUpdateTimePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateProviderUpdateTimePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateProviderUpdateTimePrimitiveTypeHolder", string(data)}, " ")
}

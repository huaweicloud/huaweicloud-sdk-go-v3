package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateProviderCreateTimePrimitiveTypeHolder struct {

	// 私有provider（private-provider）的生成时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o PrivateProviderCreateTimePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateProviderCreateTimePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateProviderCreateTimePrimitiveTypeHolder", string(data)}, " ")
}

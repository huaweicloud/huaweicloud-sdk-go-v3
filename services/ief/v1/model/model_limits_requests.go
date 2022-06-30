package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源配置限制
type LimitsRequests struct {
}

func (o LimitsRequests) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LimitsRequests struct{}"
	}

	return strings.Join([]string{"LimitsRequests", string(data)}, " ")
}

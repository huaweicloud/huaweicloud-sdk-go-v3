package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimezoneOffset 时区偏移量（仅启动类型为period时有值，单位：分钟）
type TimezoneOffset struct {
}

func (o TimezoneOffset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimezoneOffset struct{}"
	}

	return strings.Join([]string{"TimezoneOffset", string(data)}, " ")
}

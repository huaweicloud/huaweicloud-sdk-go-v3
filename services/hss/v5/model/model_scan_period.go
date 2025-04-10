package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanPeriod 启动类型，包含如下:   - day ：每天   - week : 每周   - month : 每月
type ScanPeriod struct {
}

func (o ScanPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanPeriod struct{}"
	}

	return strings.Join([]string{"ScanPeriod", string(data)}, " ")
}

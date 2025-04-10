package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanHour 扫描小时数（仅启动类型为period时有值）
type ScanHour struct {
}

func (o ScanHour) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanHour struct{}"
	}

	return strings.Join([]string{"ScanHour", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanTime 扫描时间戳，毫秒（仅启动类型为later时有值）
type ScanTime struct {
}

func (o ScanTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanTime struct{}"
	}

	return strings.Join([]string{"ScanTime", string(data)}, " ")
}

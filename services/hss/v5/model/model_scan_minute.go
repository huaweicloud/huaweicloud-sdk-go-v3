package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanMinute 扫描分钟数（仅启动类型为period时有值）
type ScanMinute struct {
}

func (o ScanMinute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanMinute struct{}"
	}

	return strings.Join([]string{"ScanMinute", string(data)}, " ")
}

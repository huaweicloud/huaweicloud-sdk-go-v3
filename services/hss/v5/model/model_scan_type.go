package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanType 任务类型，包含如下:   - quick ：快速扫描   - full : 全盘扫描   - custom : 自定义扫描
type ScanType struct {
}

func (o ScanType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanType struct{}"
	}

	return strings.Join([]string{"ScanType", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanProgress **参数解释**： 扫描进度 **取值范围**： 字符串格式，支持百分比（如“50%”）或0-100的数值字符串
type ScanProgress struct {
}

func (o ScanProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanProgress struct{}"
	}

	return strings.Join([]string{"ScanProgress", string(data)}, " ")
}

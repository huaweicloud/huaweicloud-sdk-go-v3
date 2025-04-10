package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanProgress 扫描进度
type ScanProgress struct {
}

func (o ScanProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanProgress struct{}"
	}

	return strings.Join([]string{"ScanProgress", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanDir 扫描目录，多个用;分隔
type ScanDir struct {
}

func (o ScanDir) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanDir struct{}"
	}

	return strings.Join([]string{"ScanDir", string(data)}, " ")
}

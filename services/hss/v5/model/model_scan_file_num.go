package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanFileNum **参数解释**: 已扫描的文件数量 **取值范围**: 非负整数，最小值0；单位：个
type ScanFileNum struct {
}

func (o ScanFileNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanFileNum struct{}"
	}

	return strings.Join([]string{"ScanFileNum", string(data)}, " ")
}

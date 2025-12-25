package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcRdsVolume Volume信息。
type HwcRdsVolume struct {

	// 磁盘类型。
	Type string `json:"type"`

	// 磁盘大小。
	Size int32 `json:"size"`
}

func (o HwcRdsVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcRdsVolume struct{}"
	}

	return strings.Join([]string{"HwcRdsVolume", string(data)}, " ")
}

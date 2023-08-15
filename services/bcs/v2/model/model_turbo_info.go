package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TurboInfo 极速文件存储卷信息
type TurboInfo struct {

	// 共享方式，固定值为“STANDARD”
	ShareType string `json:"share_type"`

	// 类型，固定值为“efs-ha”
	Type string `json:"type"`

	// 可用区，可填空字符串(\"\")。
	AvailableZone string `json:"available_zone"`

	// 规格，固定值为“sfs.turbo.standard”
	ResourceSpecCode string `json:"resource_spec_code"`
}

func (o TurboInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TurboInfo struct{}"
	}

	return strings.Join([]string{"TurboInfo", string(data)}, " ")
}

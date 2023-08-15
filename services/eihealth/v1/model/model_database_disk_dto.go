package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseDiskDto struct {

	// 磁盘类型
	Type string `json:"type"`

	// 磁盘大小，单位GB
	Space int32 `json:"space"`

	// 是否加密
	Encrypt bool `json:"encrypt"`

	// 磁盘已使用量，单位GB
	Used float64 `json:"used"`
}

func (o DatabaseDiskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseDiskDto struct{}"
	}

	return strings.Join([]string{"DatabaseDiskDto", string(data)}, " ")
}

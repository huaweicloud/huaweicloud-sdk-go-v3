package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PackageInfo 流水线产物
type PackageInfo struct {

	// 产物名
	Name *string `json:"name,omitempty"`

	// 产物类型
	PackageType *string `json:"packageType,omitempty"`

	// 产物版本号
	Version *string `json:"version,omitempty"`

	// 产物下载地址
	DownloadUrl *string `json:"downloadUrl,omitempty"`
}

func (o PackageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageInfo struct{}"
	}

	return strings.Join([]string{"PackageInfo", string(data)}, " ")
}

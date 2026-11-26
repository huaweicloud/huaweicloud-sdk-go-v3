package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Platform struct {

	// 制品适用的CPU架构。
	Architecture *string `json:"architecture,omitempty"`

	// 制品适用的操作系统。
	Os *string `json:"os,omitempty"`

	// 制品适用的操作系统版本。
	OsVersion *string `json:"os.version,omitempty"`

	// 制品适用的操作系统特性。
	OsFeatures *[]string `json:"os.features,omitempty"`

	// 同一系列 CPU 下的细分型号，例如 Intel i5-13400、i5-13600 属于同代不同 variant。
	Variant *string `json:"variant,omitempty"`
}

func (o Platform) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Platform struct{}"
	}

	return strings.Join([]string{"Platform", string(data)}, " ")
}

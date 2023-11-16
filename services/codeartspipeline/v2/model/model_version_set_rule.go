package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VersionSetRule 依赖规则列表
type VersionSetRule struct {

	// 是否开启
	Enable *bool `json:"enable,omitempty"`

	// 依赖类型
	Ecosystem *string `json:"ecosystem,omitempty"`

	// 包名称
	PackageName *string `json:"package_name,omitempty"`

	// 包版本
	PackageVersion *string `json:"package_version,omitempty"`

	// 规则说明
	Description *string `json:"description,omitempty"`

	// 比较规则
	Predicate *string `json:"predicate,omitempty"`
}

func (o VersionSetRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionSetRule struct{}"
	}

	return strings.Join([]string{"VersionSetRule", string(data)}, " ")
}

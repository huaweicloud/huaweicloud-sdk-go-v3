package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PackageUsage 套餐用量信息
type PackageUsage struct {

	// 套餐类型
	Name *string `json:"name,omitempty"`

	// 套餐用量
	UsedPercent *int32 `json:"used_percent,omitempty"`
}

func (o PackageUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageUsage struct{}"
	}

	return strings.Join([]string{"PackageUsage", string(data)}, " ")
}

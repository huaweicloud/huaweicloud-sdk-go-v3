package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 获取租户的习题库过滤字段
type PackageFilter struct {

	// 需查询的习题库名称
	Name *string `json:"name,omitempty"`

	// 标签名称列表
	TagNames *[]string `json:"tag_names,omitempty"`
}

func (o PackageFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageFilter struct{}"
	}

	return strings.Join([]string{"PackageFilter", string(data)}, " ")
}

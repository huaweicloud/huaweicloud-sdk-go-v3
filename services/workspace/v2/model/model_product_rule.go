package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductRule 产品规则信息,进程名称、产品名称、发布者的各项都需要匹配，为空或者*表示任意匹配。 例如： 发布者：* 产品名称：Huawei 进程名称：* 只要产品名称是Huawei的就可以认为是匹配的。
type ProductRule struct {

	// 识别条件。
	IdentifyCondition *string `json:"identify_condition,omitempty"`

	// 发布者名称： 1. 名称允许可见字符或空格，不可为全空格。 2. 长度0~1024个字符。 3. 为空或者*号表示任意匹配。
	Publisher *string `json:"publisher,omitempty"`

	// 产品名称： 1. 名称允许可见字符或空格，不可为全空格。 2. 长度0~128个字符。 3. 为空或者*号表示任意匹配。
	ProductName *string `json:"product_name,omitempty"`

	// 进程名称： 1. 名称允许可见字符或空格，不可为全空格。 2. 长度0~128个字符。 3. 为空或者*号表示任意匹配。
	ProcessName *string `json:"process_name,omitempty"`

	// 操作系统类型，仅内置规则存在该字段。 1. 名称允许可见字符或空格，不可为全空格。 2. 长度0~128个字符。
	SupportOs *string `json:"support_os,omitempty"`

	// 版本号： 1. 名称允许可见字符或空格，不可为全空格。 2. 长度0~128个字符。 3. 为空或者*号表示任意匹配。
	Version *string `json:"version,omitempty"`

	// 产品版本号： 1. 允许可见字符。 2. 长度0~128个字符。
	ProductVersion *string `json:"product_version,omitempty"`
}

func (o ProductRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductRule struct{}"
	}

	return strings.Join([]string{"ProductRule", string(data)}, " ")
}

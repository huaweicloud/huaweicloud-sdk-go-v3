package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableZone struct {

	// 可用区名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 实例创建失败错误信息
	Id *string `json:"id,omitempty" xml:"id"`

	// 可用区编码。
	Code *string `json:"code,omitempty" xml:"code"`

	// 可用区端口号。
	Port *string `json:"port,omitempty" xml:"port"`

	LocalName *LocalName `json:"local_name,omitempty" xml:"local_name"`

	// 可用区支持的实例规格。
	Specs map[string]bool `json:"specs,omitempty" xml:"specs"`
}

func (o AvailableZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableZone struct{}"
	}

	return strings.Join([]string{"AvailableZone", string(data)}, " ")
}

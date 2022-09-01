package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupportAzsInfo struct {

	// 可用区编码。
	Code *string `json:"code,omitempty" xml:"code"`

	// 可用区名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 是否支持。
	Favored *bool `json:"favored,omitempty" xml:"favored"`
}

func (o SupportAzsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportAzsInfo struct{}"
	}

	return strings.Join([]string{"SupportAzsInfo", string(data)}, " ")
}

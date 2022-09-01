package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 插件模板详情-response结构体
type AddonTemplate struct {

	// API类型，固定值“Addon”，该值不可修改。
	Kind string `json:"kind" xml:"kind"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion string `json:"apiVersion" xml:"apiVersion"`

	Metadata *Metadata `json:"metadata" xml:"metadata"`

	Spec *Templatespec `json:"spec" xml:"spec"`
}

func (o AddonTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonTemplate struct{}"
	}

	return strings.Join([]string{"AddonTemplate", string(data)}, " ")
}

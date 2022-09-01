package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomTemplates struct {

	// 配置文件id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 配置文件名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 配置文件内容。
	ConfContent *string `json:"confContent,omitempty" xml:"confContent"`

	// 描述。
	Desc *string `json:"desc,omitempty" xml:"desc"`
}

func (o CustomTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomTemplates struct{}"
	}

	return strings.Join([]string{"CustomTemplates", string(data)}, " ")
}

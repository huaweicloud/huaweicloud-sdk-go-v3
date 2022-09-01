package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CatalogTargetParameters struct {

	// 目标参数名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 目标参数名称展示说明
	Label *string `json:"label,omitempty" xml:"label"`

	// 参数展示元数据，比如是否必选，输入框类型等等
	Metadata *interface{} `json:"metadata,omitempty" xml:"metadata"`
}

func (o CatalogTargetParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogTargetParameters struct{}"
	}

	return strings.Join([]string{"CatalogTargetParameters", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateProxyConfigurationItem struct {

	// 参数名。
	Name string `json:"name"`

	// 参数值。
	Value string `json:"value"`

	// 父标签类型。
	ElemType string `json:"elem_type"`
}

func (o UpdateProxyConfigurationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyConfigurationItem struct{}"
	}

	return strings.Join([]string{"UpdateProxyConfigurationItem", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Property struct {

	// 属性名称
	Name string `json:"name"`

	// 属性值
	Value string `json:"value"`
}

func (o Property) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Property struct{}"
	}

	return strings.Join([]string{"Property", string(data)}, " ")
}

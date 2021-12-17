package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 属性定义为输入参数
type InputParam struct {
	// 参数名称

	Name string `json:"name"`
	// 属性名称

	PropertyName string `json:"property_name"`
}

func (o InputParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputParam struct{}"
	}

	return strings.Join([]string{"InputParam", string(data)}, " ")
}

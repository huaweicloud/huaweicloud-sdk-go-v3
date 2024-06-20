package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorAttributeInfo struct {

	// 属性编码
	Code string `json:"code"`

	// 属性值
	Value string `json:"value"`
}

func (o FlavorAttributeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorAttributeInfo struct{}"
	}

	return strings.Join([]string{"FlavorAttributeInfo", string(data)}, " ")
}

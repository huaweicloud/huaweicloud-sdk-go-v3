package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AzCodeInfo struct {

	// |参数名称：可用区编码| |参数的约束及描述：该参数非必填，且只允许字符串|
	AzCode *string `json:"az_code,omitempty"`
}

func (o AzCodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzCodeInfo struct{}"
	}

	return strings.Join([]string{"AzCodeInfo", string(data)}, " ")
}

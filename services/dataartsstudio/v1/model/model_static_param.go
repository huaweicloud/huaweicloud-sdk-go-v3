package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StaticParam 静态参数
type StaticParam struct {

	// 静态参数名
	ParaName *string `json:"para_name,omitempty"`

	// 静态参数值
	ParaValue *string `json:"para_value,omitempty"`
}

func (o StaticParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StaticParam struct{}"
	}

	return strings.Join([]string{"StaticParam", string(data)}, " ")
}

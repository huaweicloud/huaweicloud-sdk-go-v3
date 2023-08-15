package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParamTypeLimits 枚举类型参数实体类
type ParamTypeLimits struct {

	// 枚举值可选参数
	Name *string `json:"name,omitempty"`
}

func (o ParamTypeLimits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamTypeLimits struct{}"
	}

	return strings.Join([]string{"ParamTypeLimits", string(data)}, " ")
}

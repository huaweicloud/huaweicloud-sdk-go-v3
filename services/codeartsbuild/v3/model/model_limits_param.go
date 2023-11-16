package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LimitsParam 枚举类参数限制
type LimitsParam struct {

	// 是否生效，默认为\"0\"，为生效状态
	Disable *string `json:"disable,omitempty"`

	// 参数展示的名字
	DisplayName *string `json:"display_name,omitempty"`

	// 参数名字
	Name *string `json:"name,omitempty"`
}

func (o LimitsParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LimitsParam struct{}"
	}

	return strings.Join([]string{"LimitsParam", string(data)}, " ")
}

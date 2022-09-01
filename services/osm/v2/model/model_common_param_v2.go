package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonParamV2 struct {

	// 参数标识
	ParamKey *string `json:"param_key,omitempty" xml:"param_key"`

	// 参数名称
	ParamName *string `json:"param_name,omitempty" xml:"param_name"`

	// 是否展示
	IsShow *int32 `json:"is_show,omitempty" xml:"is_show"`

	// 是否必填
	IsRequired *int32 `json:"is_required,omitempty" xml:"is_required"`
}

func (o CommonParamV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonParamV2 struct{}"
	}

	return strings.Join([]string{"CommonParamV2", string(data)}, " ")
}

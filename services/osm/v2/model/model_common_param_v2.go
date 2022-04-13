package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonParamV2 struct {
	// 参数标识

	ParamKey *string `json:"param_key,omitempty"`
	// 参数名称

	ParamName *string `json:"param_name,omitempty"`
	// 是否展示

	IsShow *int32 `json:"is_show,omitempty"`
	// 是否必填

	IsRequired *int32 `json:"is_required,omitempty"`
}

func (o CommonParamV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonParamV2 struct{}"
	}

	return strings.Join([]string{"CommonParamV2", string(data)}, " ")
}

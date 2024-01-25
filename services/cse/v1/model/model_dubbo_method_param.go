package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DubboMethodParam struct {

	// 参数键。
	ParamKey *string `json:"paramKey,omitempty"`

	// 参数来源。
	ParamSource *string `json:"paramSource,omitempty"`

	// 参数类型。
	ParamType *string `json:"paramType,omitempty"`
}

func (o DubboMethodParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DubboMethodParam struct{}"
	}

	return strings.Join([]string{"DubboMethodParam", string(data)}, " ")
}

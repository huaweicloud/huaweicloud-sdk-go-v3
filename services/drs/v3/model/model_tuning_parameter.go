package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TuningParameter 调优参数
type TuningParameter struct {

	// 参数名称。
	ParamName *string `json:"param_name,omitempty"`

	// 参数取值。
	ParamValue *string `json:"param_value,omitempty"`

	// 是否可用。
	Availability *string `json:"availability,omitempty"`

	// 参数取值范围。
	Range *string `json:"range,omitempty"`
}

func (o TuningParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TuningParameter struct{}"
	}

	return strings.Join([]string{"TuningParameter", string(data)}, " ")
}

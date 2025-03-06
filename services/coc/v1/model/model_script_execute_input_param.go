package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScriptExecuteInputParam struct {

	// 脚本入参的名称,同一个脚本，参数名不能重复
	ParamName string `json:"param_name"`

	// 脚本入参的值，默认必填。有引用参数（param_refer不为空）时，允许为空 1.参数长度为1-4096位 2.可以包含大写字母、小写字母、数字及特殊字符(_-/.* ?:\",=+@#\\[{]}) 3.禁止出现连续'.'
	ParamValue string `json:"param_value"`

	// 该参数已废弃，传入该参数不会生效。
	ParamOrder *int32 `json:"param_order,omitempty"`

	ParamRefer *ScriptExecuteParamReference `json:"param_refer,omitempty"`
}

func (o ScriptExecuteInputParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptExecuteInputParam struct{}"
	}

	return strings.Join([]string{"ScriptExecuteInputParam", string(data)}, " ")
}

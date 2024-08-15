package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScriptParamDefine 顺序参数，没有参数名字段 参数顺序，从1开始，不连续会报错  api层用不同的vo对象接收，管理面没有order字段。service层统一处理差异  拼接：sh xxx.sh  'aaa' 'a' 'b' 执行时值为空：sh xxx.sh  ” 'a' 'b'
type ScriptParamDefine struct {

	// 参数名称：只支持英文、数字、下划线
	ParamName string `json:"param_name"`

	// 参数默认值，默认必填，如果有参数引用，可为空
	ParamValue *string `json:"param_value,omitempty"`

	// 参数描述
	ParamDescription string `json:"param_description"`

	// 该参数已废弃，传入该参数不会生效。
	ParamOrder *int32 `json:"param_order,omitempty"`

	// 是否是敏感参数
	Sensitive bool `json:"sensitive"`
}

func (o ScriptParamDefine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptParamDefine struct{}"
	}

	return strings.Join([]string{"ScriptParamDefine", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VarsStructure 合规包模版参数。
type VarsStructure struct {

	// 参数名称。
	VarKey *string `json:"var_key,omitempty"`

	// 参数的值。
	VarValue *interface{} `json:"var_value,omitempty"`
}

func (o VarsStructure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VarsStructure struct{}"
	}

	return strings.Join([]string{"VarsStructure", string(data)}, " ")
}

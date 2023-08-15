package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExceptedActual 动态属性
type ExceptedActual struct {

	// 动态属性的初始值，最大长度512，value允许英文字母、数字、下划线、中划线、点、逗号、冒号、/、@、+、?、^、=、%、&、~、#、!、*
	Value *string `json:"value,omitempty"`
}

func (o ExceptedActual) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExceptedActual struct{}"
	}

	return strings.Join([]string{"ExceptedActual", string(data)}, " ")
}

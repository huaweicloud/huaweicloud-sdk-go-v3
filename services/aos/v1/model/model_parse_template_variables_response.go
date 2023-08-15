package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseTemplateVariablesResponse Response Object
type ParseTemplateVariablesResponse struct {

	// 解析模板得到的参数列表
	Variables      *[]VariableResponse `json:"variables,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ParseTemplateVariablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseTemplateVariablesResponse struct{}"
	}

	return strings.Join([]string{"ParseTemplateVariablesResponse", string(data)}, " ")
}

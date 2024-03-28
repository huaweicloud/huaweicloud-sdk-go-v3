package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGlobalVariableRequestBody struct {

	// 变量值
	VarValue string `json:"var_value"`
}

func (o UpdateGlobalVariableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalVariableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGlobalVariableRequestBody", string(data)}, " ")
}

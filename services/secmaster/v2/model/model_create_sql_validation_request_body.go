package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSqlValidationRequestBody struct {

	// Job Script 作业脚本
	Script string `json:"script"`
}

func (o CreateSqlValidationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlValidationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSqlValidationRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecuteResultWithoutKeyNoRetryRequestBody 查询SQL执行结果请求体
type ShowExecuteResultWithoutKeyNoRetryRequestBody struct {

	// SQL执行ID
	ExecuteId string `json:"execute_id"`
}

func (o ShowExecuteResultWithoutKeyNoRetryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecuteResultWithoutKeyNoRetryRequestBody struct{}"
	}

	return strings.Join([]string{"ShowExecuteResultWithoutKeyNoRetryRequestBody", string(data)}, " ")
}

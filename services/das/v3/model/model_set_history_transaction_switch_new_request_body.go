package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetHistoryTransactionSwitchNewRequestBody 设置历史事务开关请求体
type SetHistoryTransactionSwitchNewRequestBody struct {

	// 开关状态
	SwitchOn bool `json:"switch_on"`

	// 数据库类型
	EngineType string `json:"engine_type"`
}

func (o SetHistoryTransactionSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetHistoryTransactionSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetHistoryTransactionSwitchNewRequestBody", string(data)}, " ")
}

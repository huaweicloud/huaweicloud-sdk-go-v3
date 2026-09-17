package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetLongHistoryTransactionSwitchNewRequestBody 设置长历史事务开关请求体
type SetLongHistoryTransactionSwitchNewRequestBody struct {

	// 开关状态
	SwitchOn bool `json:"switch_on"`

	// 数据库类型
	EngineType string `json:"engine_type"`

	// 阈值
	Threshold int64 `json:"threshold"`
}

func (o SetLongHistoryTransactionSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetLongHistoryTransactionSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetLongHistoryTransactionSwitchNewRequestBody", string(data)}, " ")
}

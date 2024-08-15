package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScriptExecuteModel struct {
	ExecuteParam *ScriptExecuteParam `json:"execute_param"`

	// 目标实例分批信息
	ExecuteBatches []ExecuteInstancesBatchInfo `json:"execute_batches"`
}

func (o ScriptExecuteModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptExecuteModel struct{}"
	}

	return strings.Join([]string{"ScriptExecuteModel", string(data)}, " ")
}

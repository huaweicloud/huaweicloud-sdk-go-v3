package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScriptBatchDetailModel 执行批次详情
type JobScriptBatchDetailModel struct {

	// 批次索引，从1开始
	BatchIndex *int32 `json:"batch_index,omitempty"`

	// 批次内执行实例数量
	TotalInstances *int32 `json:"total_instances,omitempty"`

	// 执行实例列表，分页
	ExecuteInstances *[]ExectionInstanceModel `json:"execute_instances,omitempty"`
}

func (o JobScriptBatchDetailModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptBatchDetailModel struct{}"
	}

	return strings.Join([]string{"JobScriptBatchDetailModel", string(data)}, " ")
}

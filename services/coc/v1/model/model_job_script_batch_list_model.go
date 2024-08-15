package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobScriptBatchListModel 执行批次列表
type JobScriptBatchListModel struct {

	// 批次索引，从1开始
	BatchIndex *int32 `json:"batch_index,omitempty"`

	// 批次内实例节点数量
	TotalInstances *int32 `json:"total_instances,omitempty"`

	// 暂停继续策略
	RotationStrategy *JobScriptBatchListModelRotationStrategy `json:"rotation_strategy,omitempty"`

	// 批次标签：
	Properties *string `json:"properties,omitempty"`
}

func (o JobScriptBatchListModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptBatchListModel struct{}"
	}

	return strings.Join([]string{"JobScriptBatchListModel", string(data)}, " ")
}

type JobScriptBatchListModelRotationStrategy struct {
	value string
}

type JobScriptBatchListModelRotationStrategyEnum struct {
	CONTINUE JobScriptBatchListModelRotationStrategy
	PAUSE    JobScriptBatchListModelRotationStrategy
}

func GetJobScriptBatchListModelRotationStrategyEnum() JobScriptBatchListModelRotationStrategyEnum {
	return JobScriptBatchListModelRotationStrategyEnum{
		CONTINUE: JobScriptBatchListModelRotationStrategy{
			value: "CONTINUE",
		},
		PAUSE: JobScriptBatchListModelRotationStrategy{
			value: "PAUSE",
		},
	}
}

func (c JobScriptBatchListModelRotationStrategy) Value() string {
	return c.value
}

func (c JobScriptBatchListModelRotationStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobScriptBatchListModelRotationStrategy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

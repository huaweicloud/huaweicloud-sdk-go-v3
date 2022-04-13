package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type PipelineDeleteStatus struct {
	// 流水线名称

	Name *string `json:"name,omitempty"`
	// 流水线删除状态,deleted:删除成功,failed:删除失败,going:正在删除中

	Status *PipelineDeleteStatusStatus `json:"status,omitempty"`
}

func (o PipelineDeleteStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineDeleteStatus struct{}"
	}

	return strings.Join([]string{"PipelineDeleteStatus", string(data)}, " ")
}

type PipelineDeleteStatusStatus struct {
	value string
}

type PipelineDeleteStatusStatusEnum struct {
	DELETED PipelineDeleteStatusStatus
	FAILED  PipelineDeleteStatusStatus
	GOING   PipelineDeleteStatusStatus
}

func GetPipelineDeleteStatusStatusEnum() PipelineDeleteStatusStatusEnum {
	return PipelineDeleteStatusStatusEnum{
		DELETED: PipelineDeleteStatusStatus{
			value: "deleted",
		},
		FAILED: PipelineDeleteStatusStatus{
			value: "failed",
		},
		GOING: PipelineDeleteStatusStatus{
			value: "going",
		},
	}
}

func (c PipelineDeleteStatusStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipelineDeleteStatusStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

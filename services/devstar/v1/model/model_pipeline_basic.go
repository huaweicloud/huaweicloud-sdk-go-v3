package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PipelineBasic 流水线资源信息
type PipelineBasic struct {

	// DevStar系统生成的流水线UUID
	Uuid *string `json:"uuid,omitempty"`

	// CloudPipeline系统对应流水线ID
	Id *string `json:"id,omitempty"`

	// 流水线名称
	Name *string `json:"name,omitempty"`

	// 流水线地址
	Url *string `json:"url,omitempty"`

	// 流水线最后一次运行状态,success:成功,failed:失败,running:运行中
	LastRunningStatus *PipelineBasicLastRunningStatus `json:"last_running_status,omitempty"`
}

func (o PipelineBasic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineBasic struct{}"
	}

	return strings.Join([]string{"PipelineBasic", string(data)}, " ")
}

type PipelineBasicLastRunningStatus struct {
	value string
}

type PipelineBasicLastRunningStatusEnum struct {
	SUCCESS PipelineBasicLastRunningStatus
	FAILED  PipelineBasicLastRunningStatus
	RUNNING PipelineBasicLastRunningStatus
}

func GetPipelineBasicLastRunningStatusEnum() PipelineBasicLastRunningStatusEnum {
	return PipelineBasicLastRunningStatusEnum{
		SUCCESS: PipelineBasicLastRunningStatus{
			value: "success",
		},
		FAILED: PipelineBasicLastRunningStatus{
			value: "failed",
		},
		RUNNING: PipelineBasicLastRunningStatus{
			value: "running",
		},
	}
}

func (c PipelineBasicLastRunningStatus) Value() string {
	return c.value
}

func (c PipelineBasicLastRunningStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipelineBasicLastRunningStatus) UnmarshalJSON(b []byte) error {
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

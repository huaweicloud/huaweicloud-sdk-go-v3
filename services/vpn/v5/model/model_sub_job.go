package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// SubJob 子任务
type SubJob struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务类型
	JobType *SubJobJobType `json:"job_type,omitempty"`

	// 任务状态
	Status *SubJobStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 完成时间
	FinishedAt *sdktime.SdkTime `json:"finished_at,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o SubJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJob struct{}"
	}

	return strings.Join([]string{"SubJob", string(data)}, " ")
}

type SubJobJobType struct {
	value string
}

type SubJobJobTypeEnum struct {
	PREPARE_RESOURCE SubJobJobType
	UPGRADE_WORKER_1 SubJobJobType
	UPGRADE_WORKER_2 SubJobJobType
}

func GetSubJobJobTypeEnum() SubJobJobTypeEnum {
	return SubJobJobTypeEnum{
		PREPARE_RESOURCE: SubJobJobType{
			value: "prepare_resource",
		},
		UPGRADE_WORKER_1: SubJobJobType{
			value: "upgrade_worker_1",
		},
		UPGRADE_WORKER_2: SubJobJobType{
			value: "upgrade_worker_2",
		},
	}
}

func (c SubJobJobType) Value() string {
	return c.value
}

func (c SubJobJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubJobJobType) UnmarshalJSON(b []byte) error {
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

type SubJobStatus struct {
	value string
}

type SubJobStatusEnum struct {
	INIT      SubJobStatus
	UPGRADING SubJobStatus
	SUCCESS   SubJobStatus
	FAIL      SubJobStatus
}

func GetSubJobStatusEnum() SubJobStatusEnum {
	return SubJobStatusEnum{
		INIT: SubJobStatus{
			value: "init",
		},
		UPGRADING: SubJobStatus{
			value: "upgrading",
		},
		SUCCESS: SubJobStatus{
			value: "success",
		},
		FAIL: SubJobStatus{
			value: "fail",
		},
	}
}

func (c SubJobStatus) Value() string {
	return c.value
}

func (c SubJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubJobStatus) UnmarshalJSON(b []byte) error {
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

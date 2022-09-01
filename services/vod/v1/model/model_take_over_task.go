package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TakeOverTask struct {

	// 桶名。
	Bucket *string `json:"bucket,omitempty" xml:"bucket"`

	// 目录/文件名。
	Object *string `json:"object,omitempty" xml:"object"`

	// 托管类型。  取值如下： - 0：表示存储到点播桶 - 1：表示存储在租户桶 - 2：表示存储到租户OBS桶中，且输出目录与源文件的存储目录相同。
	HostType *int32 `json:"host_type,omitempty" xml:"host_type"`

	// 输出桶 。
	OutputBucket *string `json:"output_bucket,omitempty" xml:"output_bucket"`

	// 输出路径 。
	OutputPath *string `json:"output_path,omitempty" xml:"output_path"`

	// 任务ID。
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 托管文件类型。
	Suffix *[]string `json:"suffix,omitempty" xml:"suffix"`

	// 转码模板组 。
	TemplateGroupName *string `json:"template_group_name,omitempty" xml:"template_group_name"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 任务状态。
	Status *TakeOverTaskStatus `json:"status,omitempty" xml:"status"`

	// 媒资的任务执行描述汇总。
	ExecDesc *string `json:"exec_desc,omitempty" xml:"exec_desc"`
}

func (o TakeOverTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TakeOverTask struct{}"
	}

	return strings.Join([]string{"TakeOverTask", string(data)}, " ")
}

type TakeOverTaskStatus struct {
	value string
}

type TakeOverTaskStatusEnum struct {
	PROCESSING TakeOverTaskStatus
	SUCCEED    TakeOverTaskStatus
	FAILED     TakeOverTaskStatus
}

func GetTakeOverTaskStatusEnum() TakeOverTaskStatusEnum {
	return TakeOverTaskStatusEnum{
		PROCESSING: TakeOverTaskStatus{
			value: "PROCESSING",
		},
		SUCCEED: TakeOverTaskStatus{
			value: "SUCCEED",
		},
		FAILED: TakeOverTaskStatus{
			value: "FAILED",
		},
	}
}

func (c TakeOverTaskStatus) Value() string {
	return c.value
}

func (c TakeOverTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TakeOverTaskStatus) UnmarshalJSON(b []byte) error {
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

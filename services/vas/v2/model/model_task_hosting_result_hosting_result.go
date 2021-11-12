package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// hosting结果文件的相关信息
type TaskHostingResultHostingResult struct {
	// 结果文件result.json的过期时间

	OverdueDate *sdktime.SdkTime `json:"overdue_date,omitempty"`
	// 结果文件result.json的状态

	Status *TaskHostingResultHostingResultStatus `json:"status,omitempty"`
	// 结果文件result.json的具体内容

	Data *string `json:"data,omitempty"`
	// 结果文件result.json的大小

	FileSize *string `json:"file_size,omitempty"`
}

func (o TaskHostingResultHostingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskHostingResultHostingResult struct{}"
	}

	return strings.Join([]string{"TaskHostingResultHostingResult", string(data)}, " ")
}

type TaskHostingResultHostingResultStatus struct {
	value string
}

type TaskHostingResultHostingResultStatusEnum struct {
	NOT_GENERATED      TaskHostingResultHostingResultStatus
	AVAILABLE          TaskHostingResultHostingResultStatus
	EXCEED_IN_SIZE     TaskHostingResultHostingResultStatus
	OVERDUE            TaskHostingResultHostingResultStatus
	DELETED_MISTAKENLY TaskHostingResultHostingResultStatus
}

func GetTaskHostingResultHostingResultStatusEnum() TaskHostingResultHostingResultStatusEnum {
	return TaskHostingResultHostingResultStatusEnum{
		NOT_GENERATED: TaskHostingResultHostingResultStatus{
			value: "NOT_GENERATED",
		},
		AVAILABLE: TaskHostingResultHostingResultStatus{
			value: "AVAILABLE",
		},
		EXCEED_IN_SIZE: TaskHostingResultHostingResultStatus{
			value: "EXCEED_IN_SIZE",
		},
		OVERDUE: TaskHostingResultHostingResultStatus{
			value: "OVERDUE",
		},
		DELETED_MISTAKENLY: TaskHostingResultHostingResultStatus{
			value: "DELETED_MISTAKENLY",
		},
	}
}

func (c TaskHostingResultHostingResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskHostingResultHostingResultStatus) UnmarshalJSON(b []byte) error {
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

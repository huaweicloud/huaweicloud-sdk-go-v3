package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportTaskItem 导出任务item。
type ExportTaskItem struct {

	// 导出任务id。
	TaskId *string `json:"task_id,omitempty"`

	// 导出文件名。
	FileName *string `json:"file_name,omitempty"`

	// 导出任务的状态，取值为 CREATING, SUCCESS, FAIL, EXPIRED; CREATING为进行中，SUCCESS为成功，FAIL为失败，EXPIRED为过期。
	Status *ExportTaskItemStatus `json:"status,omitempty"`

	// 任务失败的原因。
	FailReason *string `json:"fail_reason,omitempty"`

	// 任务失败错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 导出任务的开始时间, utc时间：yyyy-MM-dd'T'HH:mm:ss.SSSZ。
	CreateTime *string `json:"create_time,omitempty"`

	// 导出任务的结束时间, utc时间：yyyy-MM-dd'T'HH:mm:ss.SSSZ。
	EndTime *string `json:"end_time,omitempty"`

	// 是否已下载。
	IsDownload *bool `json:"is_download,omitempty"`
}

func (o ExportTaskItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTaskItem struct{}"
	}

	return strings.Join([]string{"ExportTaskItem", string(data)}, " ")
}

type ExportTaskItemStatus struct {
	value string
}

type ExportTaskItemStatusEnum struct {
	CREATING ExportTaskItemStatus
	SUCCESS  ExportTaskItemStatus
	FAIL     ExportTaskItemStatus
	EXPIRED  ExportTaskItemStatus
}

func GetExportTaskItemStatusEnum() ExportTaskItemStatusEnum {
	return ExportTaskItemStatusEnum{
		CREATING: ExportTaskItemStatus{
			value: "CREATING",
		},
		SUCCESS: ExportTaskItemStatus{
			value: "SUCCESS",
		},
		FAIL: ExportTaskItemStatus{
			value: "FAIL",
		},
		EXPIRED: ExportTaskItemStatus{
			value: "EXPIRED",
		},
	}
}

func (c ExportTaskItemStatus) Value() string {
	return c.value
}

func (c ExportTaskItemStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTaskItemStatus) UnmarshalJSON(b []byte) error {
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

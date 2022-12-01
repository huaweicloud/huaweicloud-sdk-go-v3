package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 对象级对比任务信息体。
type ObjectsCompareTaskInfo struct {

	// 对比任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 对比结果。
	CompareResults *[]ObjectsCompareOverviewInfo `json:"compare_results,omitempty"`

	// 对比任务开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 对比任务状态。取值： - RUNNING：运行中。 - WAITING_FOR_RUNNING：等待启动中。 - SUCCESSFUL：完成。 - FAILED：失败。 - CANCELLED：已取消。
	Status *ObjectsCompareTaskInfoStatus `json:"status,omitempty"`

	// 导出比对结果状态。
	ExportStatus *string `json:"export_status,omitempty"`

	// 导出比对结果有效期剩余时间。
	ReportRemainSeconds *int64 `json:"report_remain_seconds,omitempty"`

	// 对比任务ID。
	CompareJobId *string `json:"compare_job_id,omitempty"`

	// 失败原因。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ObjectsCompareTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectsCompareTaskInfo struct{}"
	}

	return strings.Join([]string{"ObjectsCompareTaskInfo", string(data)}, " ")
}

type ObjectsCompareTaskInfoStatus struct {
	value string
}

type ObjectsCompareTaskInfoStatusEnum struct {
	RUNNING             ObjectsCompareTaskInfoStatus
	WAITING_FOR_RUNNING ObjectsCompareTaskInfoStatus
	SUCCESSFUL          ObjectsCompareTaskInfoStatus
	FAILED              ObjectsCompareTaskInfoStatus
	CANCELLED           ObjectsCompareTaskInfoStatus
}

func GetObjectsCompareTaskInfoStatusEnum() ObjectsCompareTaskInfoStatusEnum {
	return ObjectsCompareTaskInfoStatusEnum{
		RUNNING: ObjectsCompareTaskInfoStatus{
			value: "RUNNING",
		},
		WAITING_FOR_RUNNING: ObjectsCompareTaskInfoStatus{
			value: "WAITING_FOR_RUNNING",
		},
		SUCCESSFUL: ObjectsCompareTaskInfoStatus{
			value: "SUCCESSFUL",
		},
		FAILED: ObjectsCompareTaskInfoStatus{
			value: "FAILED",
		},
		CANCELLED: ObjectsCompareTaskInfoStatus{
			value: "CANCELLED",
		},
	}
}

func (c ObjectsCompareTaskInfoStatus) Value() string {
	return c.value
}

func (c ObjectsCompareTaskInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectsCompareTaskInfoStatus) UnmarshalJSON(b []byte) error {
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

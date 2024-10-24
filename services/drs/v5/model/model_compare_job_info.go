package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CompareJobInfo 对比列表信息体。
type CompareJobInfo struct {

	// 对比任务ID。
	Id *string `json:"id,omitempty"`

	// 对比类型。
	Type *string `json:"type,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 对比任务的状态。取值： - RUNNING：运行中。 - WAITING_FOR_RUNNING：等待启动中。 - SUCCESSFUL：完成。 - FAILED：失败。 - CANCELLED：已取消。 - TIMEOUT_INTERRUPT：超时中断。 - FULL_DOING：全量校验中。 - INCRE_DOING：增量校验中。
	Status *CompareJobInfoStatus `json:"status,omitempty"`

	// 对比计算资源。
	ComputeType *string `json:"compute_type,omitempty"`

	// 导出比对结果状态。
	ExportStatus *string `json:"export_status,omitempty"`

	// 导出比对结果有效期剩余时间。
	ReportRemainSeconds *string `json:"report_remain_seconds,omitempty"`

	// 对比任务的标签。
	CompareJobTag map[string]string `json:"compare_job_tag,omitempty"`

	// 对比任务选项。
	Options map[string]string `json:"options,omitempty"`

	// 失败原因。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 动态比对时延。
	DynamicCompareDelay *int64 `json:"dynamic_compare_delay,omitempty"`
}

func (o CompareJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareJobInfo struct{}"
	}

	return strings.Join([]string{"CompareJobInfo", string(data)}, " ")
}

type CompareJobInfoStatus struct {
	value string
}

type CompareJobInfoStatusEnum struct {
	RUNNING             CompareJobInfoStatus
	WAITING_FOR_RUNNING CompareJobInfoStatus
	SUCCESSFUL          CompareJobInfoStatus
	FAILED              CompareJobInfoStatus
	CANCELLED           CompareJobInfoStatus
	TIMEOUT_INTERRUPT   CompareJobInfoStatus
	FULL_DOING          CompareJobInfoStatus
	INCRE_DOING         CompareJobInfoStatus
}

func GetCompareJobInfoStatusEnum() CompareJobInfoStatusEnum {
	return CompareJobInfoStatusEnum{
		RUNNING: CompareJobInfoStatus{
			value: "RUNNING",
		},
		WAITING_FOR_RUNNING: CompareJobInfoStatus{
			value: "WAITING_FOR_RUNNING",
		},
		SUCCESSFUL: CompareJobInfoStatus{
			value: "SUCCESSFUL",
		},
		FAILED: CompareJobInfoStatus{
			value: "FAILED",
		},
		CANCELLED: CompareJobInfoStatus{
			value: "CANCELLED",
		},
		TIMEOUT_INTERRUPT: CompareJobInfoStatus{
			value: "TIMEOUT_INTERRUPT",
		},
		FULL_DOING: CompareJobInfoStatus{
			value: "FULL_DOING",
		},
		INCRE_DOING: CompareJobInfoStatus{
			value: "INCRE_DOING",
		},
	}
}

func (c CompareJobInfoStatus) Value() string {
	return c.value
}

func (c CompareJobInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CompareJobInfoStatus) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReportStatusResponse Response Object
type ShowReportStatusResponse struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// PDF报告生成状态:   * ungenerated - 未生成   * generating - 生成中   * generated - 已生成   * failed - 生成失败
	ReportStatus   *ShowReportStatusResponseReportStatus `json:"report_status,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ShowReportStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowReportStatusResponse", string(data)}, " ")
}

type ShowReportStatusResponseReportStatus struct {
	value string
}

type ShowReportStatusResponseReportStatusEnum struct {
	UNGENERATED ShowReportStatusResponseReportStatus
	GENERATING  ShowReportStatusResponseReportStatus
	GENERATED   ShowReportStatusResponseReportStatus
	FAILED      ShowReportStatusResponseReportStatus
}

func GetShowReportStatusResponseReportStatusEnum() ShowReportStatusResponseReportStatusEnum {
	return ShowReportStatusResponseReportStatusEnum{
		UNGENERATED: ShowReportStatusResponseReportStatus{
			value: "ungenerated",
		},
		GENERATING: ShowReportStatusResponseReportStatus{
			value: "generating",
		},
		GENERATED: ShowReportStatusResponseReportStatus{
			value: "generated",
		},
		FAILED: ShowReportStatusResponseReportStatus{
			value: "failed",
		},
	}
}

func (c ShowReportStatusResponseReportStatus) Value() string {
	return c.value
}

func (c ShowReportStatusResponseReportStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReportStatusResponseReportStatus) UnmarshalJSON(b []byte) error {
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

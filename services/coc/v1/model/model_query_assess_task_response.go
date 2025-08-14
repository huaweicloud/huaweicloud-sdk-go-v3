package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryAssessTaskResponse 评估任务数据
type QueryAssessTaskResponse struct {

	// 唯一标识ID
	Id *string `json:"id,omitempty"`

	// 应用名称
	ApplicationName *string `json:"application_name,omitempty"`

	// 应用ID
	ApplicationId *string `json:"application_id,omitempty"`

	// 评估任务状态.  no_assessment 未评估 assess_finish 评估完成 assess_failed 评估失败 assessing     评估中
	Status *QueryAssessTaskResponseStatus `json:"status,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	// 评估进度
	Progress *int32 `json:"progress,omitempty"`

	Score *float64 `json:"score,omitempty"`

	// 评估报告编号ID
	AssessReportId *string `json:"assess_report_id,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 最新评估时间
	LastAssessTime *int64 `json:"last_assess_time,omitempty"`

	// 最新更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 创建人ID
	Creator *string `json:"creator,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 操作人ID
	Operator *string `json:"operator,omitempty"`

	// 操作人名称
	OperatorName *string `json:"operator_name,omitempty"`
}

func (o QueryAssessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAssessTaskResponse struct{}"
	}

	return strings.Join([]string{"QueryAssessTaskResponse", string(data)}, " ")
}

type QueryAssessTaskResponseStatus struct {
	value string
}

type QueryAssessTaskResponseStatusEnum struct {
	NO_ASSESSMENT QueryAssessTaskResponseStatus
	ASSESS_FINISH QueryAssessTaskResponseStatus
	ASSESS_FAILED QueryAssessTaskResponseStatus
	ASSESSING     QueryAssessTaskResponseStatus
}

func GetQueryAssessTaskResponseStatusEnum() QueryAssessTaskResponseStatusEnum {
	return QueryAssessTaskResponseStatusEnum{
		NO_ASSESSMENT: QueryAssessTaskResponseStatus{
			value: "no_assessment",
		},
		ASSESS_FINISH: QueryAssessTaskResponseStatus{
			value: "assess_finish",
		},
		ASSESS_FAILED: QueryAssessTaskResponseStatus{
			value: "assess_failed",
		},
		ASSESSING: QueryAssessTaskResponseStatus{
			value: "assessing",
		},
	}
}

func (c QueryAssessTaskResponseStatus) Value() string {
	return c.value
}

func (c QueryAssessTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryAssessTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

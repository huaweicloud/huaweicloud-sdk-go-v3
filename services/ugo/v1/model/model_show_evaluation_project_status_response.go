package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEvaluationProjectStatusResponse Response Object
type ShowEvaluationProjectStatusResponse struct {

	// 评估项目ID。
	EvaluationProjectId *int32 `json:"evaluation_project_id,omitempty"`

	// 评估项目名称。
	EvaluationProjectName *string `json:"evaluation_project_name,omitempty"`

	// 评估项目状态。
	EvaluationProjectStatus *ShowEvaluationProjectStatusResponseEvaluationProjectStatus `json:"evaluation_project_status,omitempty"`

	ProjectStatusDetail *ProjectStatusDetail `json:"project_status_detail,omitempty"`

	// 源数据库类型。
	SourceDbType *string `json:"source_db_type,omitempty"`

	// 源数据库版本。
	SourceDbVersion *string `json:"source_db_version,omitempty"`

	// 目标数据库类型。
	TargetDbType *string `json:"target_db_type,omitempty"`

	// 目标数据库版本。
	TargetDbVersion *string `json:"target_db_version,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowEvaluationProjectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEvaluationProjectStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowEvaluationProjectStatusResponse", string(data)}, " ")
}

type ShowEvaluationProjectStatusResponseEvaluationProjectStatus struct {
	value string
}

type ShowEvaluationProjectStatusResponseEvaluationProjectStatusEnum struct {
	COMPLETED ShowEvaluationProjectStatusResponseEvaluationProjectStatus
	WAITING   ShowEvaluationProjectStatusResponseEvaluationProjectStatus
	PENDING   ShowEvaluationProjectStatusResponseEvaluationProjectStatus
	FAILED    ShowEvaluationProjectStatusResponseEvaluationProjectStatus
	STOPPED   ShowEvaluationProjectStatusResponseEvaluationProjectStatus
}

func GetShowEvaluationProjectStatusResponseEvaluationProjectStatusEnum() ShowEvaluationProjectStatusResponseEvaluationProjectStatusEnum {
	return ShowEvaluationProjectStatusResponseEvaluationProjectStatusEnum{
		COMPLETED: ShowEvaluationProjectStatusResponseEvaluationProjectStatus{
			value: "COMPLETED",
		},
		WAITING: ShowEvaluationProjectStatusResponseEvaluationProjectStatus{
			value: "WAITING",
		},
		PENDING: ShowEvaluationProjectStatusResponseEvaluationProjectStatus{
			value: "PENDING",
		},
		FAILED: ShowEvaluationProjectStatusResponseEvaluationProjectStatus{
			value: "FAILED",
		},
		STOPPED: ShowEvaluationProjectStatusResponseEvaluationProjectStatus{
			value: "STOPPED",
		},
	}
}

func (c ShowEvaluationProjectStatusResponseEvaluationProjectStatus) Value() string {
	return c.value
}

func (c ShowEvaluationProjectStatusResponseEvaluationProjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEvaluationProjectStatusResponseEvaluationProjectStatus) UnmarshalJSON(b []byte) error {
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

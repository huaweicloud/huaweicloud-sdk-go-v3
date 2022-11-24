package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ConfirmTargetDbTypeResponse struct {

	// 评估项目ID。
	EvaluationProjectId *int32 `json:"evaluation_project_id,omitempty"`

	// 评估项目名称。
	EvaluationProjectName *string `json:"evaluation_project_name,omitempty"`

	// 评估项目状态。
	EvaluationProjectStatus *ConfirmTargetDbTypeResponseEvaluationProjectStatus `json:"evaluation_project_status,omitempty"`

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

func (o ConfirmTargetDbTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmTargetDbTypeResponse struct{}"
	}

	return strings.Join([]string{"ConfirmTargetDbTypeResponse", string(data)}, " ")
}

type ConfirmTargetDbTypeResponseEvaluationProjectStatus struct {
	value string
}

type ConfirmTargetDbTypeResponseEvaluationProjectStatusEnum struct {
	COMPLETED ConfirmTargetDbTypeResponseEvaluationProjectStatus
	WAITING   ConfirmTargetDbTypeResponseEvaluationProjectStatus
	PENDING   ConfirmTargetDbTypeResponseEvaluationProjectStatus
	FAILED    ConfirmTargetDbTypeResponseEvaluationProjectStatus
	STOPPED   ConfirmTargetDbTypeResponseEvaluationProjectStatus
}

func GetConfirmTargetDbTypeResponseEvaluationProjectStatusEnum() ConfirmTargetDbTypeResponseEvaluationProjectStatusEnum {
	return ConfirmTargetDbTypeResponseEvaluationProjectStatusEnum{
		COMPLETED: ConfirmTargetDbTypeResponseEvaluationProjectStatus{
			value: "COMPLETED",
		},
		WAITING: ConfirmTargetDbTypeResponseEvaluationProjectStatus{
			value: "WAITING",
		},
		PENDING: ConfirmTargetDbTypeResponseEvaluationProjectStatus{
			value: "PENDING",
		},
		FAILED: ConfirmTargetDbTypeResponseEvaluationProjectStatus{
			value: "FAILED",
		},
		STOPPED: ConfirmTargetDbTypeResponseEvaluationProjectStatus{
			value: "STOPPED",
		},
	}
}

func (c ConfirmTargetDbTypeResponseEvaluationProjectStatus) Value() string {
	return c.value
}

func (c ConfirmTargetDbTypeResponseEvaluationProjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmTargetDbTypeResponseEvaluationProjectStatus) UnmarshalJSON(b []byte) error {
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

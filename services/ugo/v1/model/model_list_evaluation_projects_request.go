package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListEvaluationProjectsRequest struct {

	// 评估项目名称（模糊搜索）。
	EvaluationProjectName *string `json:"evaluation_project_name,omitempty"`

	// 评估项目状态。
	EvaluationProjectStatus *ListEvaluationProjectsRequestEvaluationProjectStatus `json:"evaluation_project_status,omitempty"`

	// 分页查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEvaluationProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEvaluationProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListEvaluationProjectsRequest", string(data)}, " ")
}

type ListEvaluationProjectsRequestEvaluationProjectStatus struct {
	value string
}

type ListEvaluationProjectsRequestEvaluationProjectStatusEnum struct {
	COMPLETED ListEvaluationProjectsRequestEvaluationProjectStatus
	PENDING   ListEvaluationProjectsRequestEvaluationProjectStatus
	FAILED    ListEvaluationProjectsRequestEvaluationProjectStatus
	STOPPED   ListEvaluationProjectsRequestEvaluationProjectStatus
}

func GetListEvaluationProjectsRequestEvaluationProjectStatusEnum() ListEvaluationProjectsRequestEvaluationProjectStatusEnum {
	return ListEvaluationProjectsRequestEvaluationProjectStatusEnum{
		COMPLETED: ListEvaluationProjectsRequestEvaluationProjectStatus{
			value: "COMPLETED",
		},
		PENDING: ListEvaluationProjectsRequestEvaluationProjectStatus{
			value: "PENDING",
		},
		FAILED: ListEvaluationProjectsRequestEvaluationProjectStatus{
			value: "FAILED",
		},
		STOPPED: ListEvaluationProjectsRequestEvaluationProjectStatus{
			value: "STOPPED",
		},
	}
}

func (c ListEvaluationProjectsRequestEvaluationProjectStatus) Value() string {
	return c.value
}

func (c ListEvaluationProjectsRequestEvaluationProjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEvaluationProjectsRequestEvaluationProjectStatus) UnmarshalJSON(b []byte) error {
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

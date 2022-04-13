package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListInstancesDetailsRequest struct {
	// 引擎类型：rabbitmq，参数缺失查询所有实例。

	Engine *string `json:"engine,omitempty"`
	// 实例名称。

	Name *string `json:"name,omitempty"`
	// 实例ID。

	InstanceId *string `json:"instance_id,omitempty"`
	// 实例状态。

	Status *ListInstancesDetailsRequestStatus `json:"status,omitempty"`
	// 是否返回创建失败的实例数。  当参数值为“true”时，返回创建失败的实例数。参数值为“false”或者其他值，不返回创建失败的实例数。

	IncludeFailure *ListInstancesDetailsRequestIncludeFailure `json:"include_failure,omitempty"`
	// 是否按照实例名称进行精确匹配查询。  默认为“false”，表示模糊匹配实例名称查询。若参数值为“true”表示按照实例名称进行精确匹配查询。

	ExactMatchName *ListInstancesDetailsRequestExactMatchName `json:"exact_match_name,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListInstancesDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesDetailsRequest", string(data)}, " ")
}

type ListInstancesDetailsRequestStatus struct {
	value string
}

type ListInstancesDetailsRequestStatusEnum struct {
	CREATING     ListInstancesDetailsRequestStatus
	CREATEFAILED ListInstancesDetailsRequestStatus
	RUNNING      ListInstancesDetailsRequestStatus
	ERROR        ListInstancesDetailsRequestStatus
	STARTING     ListInstancesDetailsRequestStatus
	RESTARTING   ListInstancesDetailsRequestStatus
	CLOSING      ListInstancesDetailsRequestStatus
	FROZEN       ListInstancesDetailsRequestStatus
}

func GetListInstancesDetailsRequestStatusEnum() ListInstancesDetailsRequestStatusEnum {
	return ListInstancesDetailsRequestStatusEnum{
		CREATING: ListInstancesDetailsRequestStatus{
			value: "CREATING",
		},
		CREATEFAILED: ListInstancesDetailsRequestStatus{
			value: "CREATEFAILED",
		},
		RUNNING: ListInstancesDetailsRequestStatus{
			value: "RUNNING",
		},
		ERROR: ListInstancesDetailsRequestStatus{
			value: "ERROR",
		},
		STARTING: ListInstancesDetailsRequestStatus{
			value: "STARTING",
		},
		RESTARTING: ListInstancesDetailsRequestStatus{
			value: "RESTARTING",
		},
		CLOSING: ListInstancesDetailsRequestStatus{
			value: "CLOSING",
		},
		FROZEN: ListInstancesDetailsRequestStatus{
			value: "FROZEN",
		},
	}
}

func (c ListInstancesDetailsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesDetailsRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListInstancesDetailsRequestIncludeFailure struct {
	value string
}

type ListInstancesDetailsRequestIncludeFailureEnum struct {
	TRUE  ListInstancesDetailsRequestIncludeFailure
	FALSE ListInstancesDetailsRequestIncludeFailure
}

func GetListInstancesDetailsRequestIncludeFailureEnum() ListInstancesDetailsRequestIncludeFailureEnum {
	return ListInstancesDetailsRequestIncludeFailureEnum{
		TRUE: ListInstancesDetailsRequestIncludeFailure{
			value: "true",
		},
		FALSE: ListInstancesDetailsRequestIncludeFailure{
			value: "false",
		},
	}
}

func (c ListInstancesDetailsRequestIncludeFailure) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesDetailsRequestIncludeFailure) UnmarshalJSON(b []byte) error {
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

type ListInstancesDetailsRequestExactMatchName struct {
	value string
}

type ListInstancesDetailsRequestExactMatchNameEnum struct {
	TRUE  ListInstancesDetailsRequestExactMatchName
	FALSE ListInstancesDetailsRequestExactMatchName
}

func GetListInstancesDetailsRequestExactMatchNameEnum() ListInstancesDetailsRequestExactMatchNameEnum {
	return ListInstancesDetailsRequestExactMatchNameEnum{
		TRUE: ListInstancesDetailsRequestExactMatchName{
			value: "true",
		},
		FALSE: ListInstancesDetailsRequestExactMatchName{
			value: "false",
		},
	}
}

func (c ListInstancesDetailsRequestExactMatchName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesDetailsRequestExactMatchName) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstancesDetailsRequest Request Object
type ListInstancesDetailsRequest struct {

	// 引擎类型：rabbitmq。
	Engine ListInstancesDetailsRequestEngine `json:"engine"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例状态，[详细状态说明请参考[实例状态说明](rabbitmq-api-180514012.xml)](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,sbc,hk_sbc,g42,hk_g42,tm,hk_tm)[详细状态说明请参考[实例状态说明](kafka-api-180514012.xml)](tag:hcs)。
	Status *ListInstancesDetailsRequestStatus `json:"status,omitempty"`

	// 是否返回创建失败的实例数。  当参数值为“true”时，返回创建失败的实例数。参数值为“false”或者其他值，不返回创建失败的实例数。
	IncludeFailure *ListInstancesDetailsRequestIncludeFailure `json:"include_failure,omitempty"`

	// 是否按照实例名称进行精确匹配查询。  默认为“false”，表示模糊匹配实例名称查询。若参数值为“true”表示按照实例名称进行精确匹配查询。
	ExactMatchName *ListInstancesDetailsRequestExactMatchName `json:"exact_match_name,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *string `json:"offset,omitempty"`

	// 当次查询返回的最大实例个数，默认值为10，取值范围为1~50。
	Limit *string `json:"limit,omitempty"`
}

func (o ListInstancesDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesDetailsRequest", string(data)}, " ")
}

type ListInstancesDetailsRequestEngine struct {
	value string
}

type ListInstancesDetailsRequestEngineEnum struct {
	RABBITMQ ListInstancesDetailsRequestEngine
}

func GetListInstancesDetailsRequestEngineEnum() ListInstancesDetailsRequestEngineEnum {
	return ListInstancesDetailsRequestEngineEnum{
		RABBITMQ: ListInstancesDetailsRequestEngine{
			value: "rabbitmq",
		},
	}
}

func (c ListInstancesDetailsRequestEngine) Value() string {
	return c.value
}

func (c ListInstancesDetailsRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesDetailsRequestEngine) UnmarshalJSON(b []byte) error {
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

type ListInstancesDetailsRequestStatus struct {
	value string
}

type ListInstancesDetailsRequestStatusEnum struct {
	CREATING           ListInstancesDetailsRequestStatus
	RUNNING            ListInstancesDetailsRequestStatus
	RESTARTING         ListInstancesDetailsRequestStatus
	DELETING           ListInstancesDetailsRequestStatus
	ERROR              ListInstancesDetailsRequestStatus
	CREATEFAILED       ListInstancesDetailsRequestStatus
	FREEZING           ListInstancesDetailsRequestStatus
	FROZEN             ListInstancesDetailsRequestStatus
	EXTENDING          ListInstancesDetailsRequestStatus
	SHRINKING          ListInstancesDetailsRequestStatus
	EXTENDEDFAILED     ListInstancesDetailsRequestStatus
	CONFIGURING        ListInstancesDetailsRequestStatus
	ROLLBACK           ListInstancesDetailsRequestStatus
	ROLLBACKFAILED     ListInstancesDetailsRequestStatus
	VOLUMETYPECHANGING ListInstancesDetailsRequestStatus
}

func GetListInstancesDetailsRequestStatusEnum() ListInstancesDetailsRequestStatusEnum {
	return ListInstancesDetailsRequestStatusEnum{
		CREATING: ListInstancesDetailsRequestStatus{
			value: "CREATING",
		},
		RUNNING: ListInstancesDetailsRequestStatus{
			value: "RUNNING",
		},
		RESTARTING: ListInstancesDetailsRequestStatus{
			value: "RESTARTING",
		},
		DELETING: ListInstancesDetailsRequestStatus{
			value: "DELETING",
		},
		ERROR: ListInstancesDetailsRequestStatus{
			value: "ERROR",
		},
		CREATEFAILED: ListInstancesDetailsRequestStatus{
			value: "CREATEFAILED",
		},
		FREEZING: ListInstancesDetailsRequestStatus{
			value: "FREEZING",
		},
		FROZEN: ListInstancesDetailsRequestStatus{
			value: "FROZEN",
		},
		EXTENDING: ListInstancesDetailsRequestStatus{
			value: "EXTENDING",
		},
		SHRINKING: ListInstancesDetailsRequestStatus{
			value: "SHRINKING",
		},
		EXTENDEDFAILED: ListInstancesDetailsRequestStatus{
			value: "EXTENDEDFAILED",
		},
		CONFIGURING: ListInstancesDetailsRequestStatus{
			value: "CONFIGURING",
		},
		ROLLBACK: ListInstancesDetailsRequestStatus{
			value: "ROLLBACK",
		},
		ROLLBACKFAILED: ListInstancesDetailsRequestStatus{
			value: "ROLLBACKFAILED",
		},
		VOLUMETYPECHANGING: ListInstancesDetailsRequestStatus{
			value: "VOLUMETYPECHANGING",
		},
	}
}

func (c ListInstancesDetailsRequestStatus) Value() string {
	return c.value
}

func (c ListInstancesDetailsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesDetailsRequestStatus) UnmarshalJSON(b []byte) error {
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

func (c ListInstancesDetailsRequestIncludeFailure) Value() string {
	return c.value
}

func (c ListInstancesDetailsRequestIncludeFailure) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesDetailsRequestIncludeFailure) UnmarshalJSON(b []byte) error {
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

func (c ListInstancesDetailsRequestExactMatchName) Value() string {
	return c.value
}

func (c ListInstancesDetailsRequestExactMatchName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesDetailsRequestExactMatchName) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstancesRequest Request Object
type ListInstancesRequest struct {

	// **参数解释**： 消息引擎。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Engine ListInstancesRequestEngine `json:"engine"`

	// **参数解释**： 实例名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 实例ID。获取方法如下：登录RocketMQ控制台，在RocketMQ实例详情页面查找实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 实例状态，[详细状态说明请参考[实例状态说明](hrm-api-0010.xml)。](tag:hws,hws_hk,ctc,hws_eu,ocb,g42,hk_g42,tm,sbc,hk_sbc,cmcc,hk_tm)[详细状态说明请参考[实例状态说明](kafka-api-180514012.xml)。](tag:hcs,fcs) **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Status *ListInstancesRequestStatus `json:"status,omitempty"`

	// **参数解释**： 是否返回创建失败的实例数。 **约束限制**： 不涉及。 **取值范围**： - 'true'：返回创建失败的实例数。 - 'false'：不返回创建失败的实例数。  **默认取值**： 不涉及。
	IncludeFailure *ListInstancesRequestIncludeFailure `json:"include_failure,omitempty"`

	// **参数解释**： 是否按照实例名称进行精确匹配查询。 **约束限制**： 不涉及。 **取值范围**： - 'true'：表示按照实例名称进行精确匹配查询。 - 'false'：表示模糊匹配实例名称查询。  **默认取值**： 'false'。
	ExactMatchName *ListInstancesRequestExactMatchName `json:"exact_match_name,omitempty"`

	// **参数解释**： 企业项目ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 当次查询返回的实例最大个数。 **约束限制**： 不涉及。 **取值范围**： 1~50。 **默认取值**： 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

type ListInstancesRequestEngine struct {
	value string
}

type ListInstancesRequestEngineEnum struct {
	ROCKETMQ ListInstancesRequestEngine
}

func GetListInstancesRequestEngineEnum() ListInstancesRequestEngineEnum {
	return ListInstancesRequestEngineEnum{
		ROCKETMQ: ListInstancesRequestEngine{
			value: "rocketmq",
		},
	}
}

func (c ListInstancesRequestEngine) Value() string {
	return c.value
}

func (c ListInstancesRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestEngine) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestStatus struct {
	value string
}

type ListInstancesRequestStatusEnum struct {
	CREATING           ListInstancesRequestStatus
	RUNNING            ListInstancesRequestStatus
	RESTARTING         ListInstancesRequestStatus
	DELETING           ListInstancesRequestStatus
	ERROR              ListInstancesRequestStatus
	CREATEFAILED       ListInstancesRequestStatus
	FREEZING           ListInstancesRequestStatus
	FROZEN             ListInstancesRequestStatus
	EXTENDING          ListInstancesRequestStatus
	SHRINKING          ListInstancesRequestStatus
	EXTENDEDFAILED     ListInstancesRequestStatus
	CONFIGURING        ListInstancesRequestStatus
	ROLLBACK           ListInstancesRequestStatus
	ROLLBACKFAILED     ListInstancesRequestStatus
	VOLUMETYPECHANGING ListInstancesRequestStatus
}

func GetListInstancesRequestStatusEnum() ListInstancesRequestStatusEnum {
	return ListInstancesRequestStatusEnum{
		CREATING: ListInstancesRequestStatus{
			value: "CREATING",
		},
		RUNNING: ListInstancesRequestStatus{
			value: "RUNNING",
		},
		RESTARTING: ListInstancesRequestStatus{
			value: "RESTARTING",
		},
		DELETING: ListInstancesRequestStatus{
			value: "DELETING",
		},
		ERROR: ListInstancesRequestStatus{
			value: "ERROR",
		},
		CREATEFAILED: ListInstancesRequestStatus{
			value: "CREATEFAILED",
		},
		FREEZING: ListInstancesRequestStatus{
			value: "FREEZING",
		},
		FROZEN: ListInstancesRequestStatus{
			value: "FROZEN",
		},
		EXTENDING: ListInstancesRequestStatus{
			value: "EXTENDING",
		},
		SHRINKING: ListInstancesRequestStatus{
			value: "SHRINKING",
		},
		EXTENDEDFAILED: ListInstancesRequestStatus{
			value: "EXTENDEDFAILED",
		},
		CONFIGURING: ListInstancesRequestStatus{
			value: "CONFIGURING",
		},
		ROLLBACK: ListInstancesRequestStatus{
			value: "ROLLBACK",
		},
		ROLLBACKFAILED: ListInstancesRequestStatus{
			value: "ROLLBACKFAILED",
		},
		VOLUMETYPECHANGING: ListInstancesRequestStatus{
			value: "VOLUMETYPECHANGING",
		},
	}
}

func (c ListInstancesRequestStatus) Value() string {
	return c.value
}

func (c ListInstancesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestIncludeFailure struct {
	value string
}

type ListInstancesRequestIncludeFailureEnum struct {
	TRUE  ListInstancesRequestIncludeFailure
	FALSE ListInstancesRequestIncludeFailure
}

func GetListInstancesRequestIncludeFailureEnum() ListInstancesRequestIncludeFailureEnum {
	return ListInstancesRequestIncludeFailureEnum{
		TRUE: ListInstancesRequestIncludeFailure{
			value: "true",
		},
		FALSE: ListInstancesRequestIncludeFailure{
			value: "false",
		},
	}
}

func (c ListInstancesRequestIncludeFailure) Value() string {
	return c.value
}

func (c ListInstancesRequestIncludeFailure) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestIncludeFailure) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestExactMatchName struct {
	value string
}

type ListInstancesRequestExactMatchNameEnum struct {
	TRUE  ListInstancesRequestExactMatchName
	FALSE ListInstancesRequestExactMatchName
}

func GetListInstancesRequestExactMatchNameEnum() ListInstancesRequestExactMatchNameEnum {
	return ListInstancesRequestExactMatchNameEnum{
		TRUE: ListInstancesRequestExactMatchName{
			value: "true",
		},
		FALSE: ListInstancesRequestExactMatchName{
			value: "false",
		},
	}
}

func (c ListInstancesRequestExactMatchName) Value() string {
	return c.value
}

func (c ListInstancesRequestExactMatchName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestExactMatchName) UnmarshalJSON(b []byte) error {
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

package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListAllScalingV2PoliciesRequest struct {
	ScalingResourceId *string `json:"scaling_resource_id,omitempty"`

	ScalingResourceType *ListAllScalingV2PoliciesRequestScalingResourceType `json:"scaling_resource_type,omitempty"`

	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`

	ScalingPolicyType *ListAllScalingV2PoliciesRequestScalingPolicyType `json:"scaling_policy_type,omitempty"`

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	SortBy *ListAllScalingV2PoliciesRequestSortBy `json:"sort_by,omitempty"`

	Order *ListAllScalingV2PoliciesRequestOrder `json:"order,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListAllScalingV2PoliciesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAllScalingV2PoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListAllScalingV2PoliciesRequest", string(data)}, " ")
}

type ListAllScalingV2PoliciesRequestScalingResourceType struct {
	value string
}

type ListAllScalingV2PoliciesRequestScalingResourceTypeEnum struct {
	SCALING_GROUP ListAllScalingV2PoliciesRequestScalingResourceType
	BANDWIDTH     ListAllScalingV2PoliciesRequestScalingResourceType
}

func GetListAllScalingV2PoliciesRequestScalingResourceTypeEnum() ListAllScalingV2PoliciesRequestScalingResourceTypeEnum {
	return ListAllScalingV2PoliciesRequestScalingResourceTypeEnum{
		SCALING_GROUP: ListAllScalingV2PoliciesRequestScalingResourceType{
			value: "SCALING_GROUP",
		},
		BANDWIDTH: ListAllScalingV2PoliciesRequestScalingResourceType{
			value: "BANDWIDTH",
		},
	}
}

func (c ListAllScalingV2PoliciesRequestScalingResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListAllScalingV2PoliciesRequestScalingResourceType) UnmarshalJSON(b []byte) error {
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

type ListAllScalingV2PoliciesRequestScalingPolicyType struct {
	value string
}

type ListAllScalingV2PoliciesRequestScalingPolicyTypeEnum struct {
	ALARM      ListAllScalingV2PoliciesRequestScalingPolicyType
	SCHEDULED  ListAllScalingV2PoliciesRequestScalingPolicyType
	RECURRENCE ListAllScalingV2PoliciesRequestScalingPolicyType
}

func GetListAllScalingV2PoliciesRequestScalingPolicyTypeEnum() ListAllScalingV2PoliciesRequestScalingPolicyTypeEnum {
	return ListAllScalingV2PoliciesRequestScalingPolicyTypeEnum{
		ALARM: ListAllScalingV2PoliciesRequestScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: ListAllScalingV2PoliciesRequestScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: ListAllScalingV2PoliciesRequestScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c ListAllScalingV2PoliciesRequestScalingPolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListAllScalingV2PoliciesRequestScalingPolicyType) UnmarshalJSON(b []byte) error {
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

type ListAllScalingV2PoliciesRequestSortBy struct {
	value string
}

type ListAllScalingV2PoliciesRequestSortByEnum struct {
	POLICY_NAME       ListAllScalingV2PoliciesRequestSortBy
	TRIGGER_CONDITION ListAllScalingV2PoliciesRequestSortBy
	CREATE_TIME       ListAllScalingV2PoliciesRequestSortBy
}

func GetListAllScalingV2PoliciesRequestSortByEnum() ListAllScalingV2PoliciesRequestSortByEnum {
	return ListAllScalingV2PoliciesRequestSortByEnum{
		POLICY_NAME: ListAllScalingV2PoliciesRequestSortBy{
			value: "POLICY_NAME",
		},
		TRIGGER_CONDITION: ListAllScalingV2PoliciesRequestSortBy{
			value: "TRIGGER_CONDITION",
		},
		CREATE_TIME: ListAllScalingV2PoliciesRequestSortBy{
			value: "CREATE_TIME",
		},
	}
}

func (c ListAllScalingV2PoliciesRequestSortBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListAllScalingV2PoliciesRequestSortBy) UnmarshalJSON(b []byte) error {
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

type ListAllScalingV2PoliciesRequestOrder struct {
	value string
}

type ListAllScalingV2PoliciesRequestOrderEnum struct {
	ASC  ListAllScalingV2PoliciesRequestOrder
	DESC ListAllScalingV2PoliciesRequestOrder
}

func GetListAllScalingV2PoliciesRequestOrderEnum() ListAllScalingV2PoliciesRequestOrderEnum {
	return ListAllScalingV2PoliciesRequestOrderEnum{
		ASC: ListAllScalingV2PoliciesRequestOrder{
			value: "ASC",
		},
		DESC: ListAllScalingV2PoliciesRequestOrder{
			value: "DESC",
		},
	}
}

func (c ListAllScalingV2PoliciesRequestOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListAllScalingV2PoliciesRequestOrder) UnmarshalJSON(b []byte) error {
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

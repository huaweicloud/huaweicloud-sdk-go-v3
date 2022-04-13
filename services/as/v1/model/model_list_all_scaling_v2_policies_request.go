package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListAllScalingV2PoliciesRequest struct {
	// 伸缩组ID。

	ScalingResourceId *string `json:"scaling_resource_id,omitempty"`
	// 伸缩资源类型：伸缩组：SCALING_GROUP；带宽：BANDWIDTH

	ScalingResourceType *ListAllScalingV2PoliciesRequestScalingResourceType `json:"scaling_resource_type,omitempty"`
	// 伸缩策略名称。

	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`
	// 策略类型。  告警策略：ALARM ,定时策略：SCHEDULED, 周期策略：RECURRENCE

	ScalingPolicyType *ListAllScalingV2PoliciesRequestScalingPolicyType `json:"scaling_policy_type,omitempty"`
	// 伸缩策略ID。

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`
	// 查询的起始行号，默认为0。

	StartNumber *int32 `json:"start_number,omitempty"`
	// 查询记录数，默认20，最大100。

	Limit *int32 `json:"limit,omitempty"`
	// 排序方法POLICY_NAME：根据策略名称排序;TRIGGER_CONDITION：根据触发条件排序，如升序下，告警策略最先，其余根据最近一次触发时间升序排列;CREATE_TIME：根据策略的创建时间排序。

	SortBy *ListAllScalingV2PoliciesRequestSortBy `json:"sort_by,omitempty"`
	// 排序顺序ASC：升序；DESC：降序

	Order *ListAllScalingV2PoliciesRequestOrder `json:"order,omitempty"`
	// 企业项目ID。  当scaling_resource_type指定为：SCALING_GROUP 传入all_granted_eps时：  华为云帐号和拥有全局权限的IAM用户可以查询该用户所有的伸缩组对应的伸缩策略。 授予部分企业项目的IAM用户，可以查询该用户所有授权企业项目下的伸缩组对应的伸缩策略。 说明： 如果授予部分企业项目的IAM用户拥有超过100个企业项目，则只能返回有权限的前100个企业项目对应伸缩组的伸缩策略列表。  当scaling_resource_type指定为：BANDWIDTH 传入all_granted_eps时:  华为云帐号和拥有全局权限的IAM用户可以查询该用户所有带宽对应的伸缩策略。 授予部分企业项目的IAM用户，可以查询该用户所有授权企业项目下的带宽对应的伸缩策略，带宽在all_granted_eps场景下返回策略请参见[《EIP接口参口》查询带宽列表](https://support.huaweicloud.com/api-eip/eip_apiBandwidth_0002.html)。 不指定scaling_resource_type 当传入all_granted_eps时：  华为云帐号和拥有全局权限的IAM用户可以查询该用户所有的伸缩组和带宽对应的伸缩策略。 授予部分企业项目的IAM用户，可以查询该用户所有授权企业项目下的伸缩组和带宽对应的伸缩策略。 说明： 如果授予部分企业项目的IAM用户拥有超过100个企业项目，则只能返回有权限的前100个企业项目对应伸缩组的伸缩策略列表；带宽在all_granted_eps场景下返回策略请参见[《EIP接口参口》查询带宽列表](https://support.huaweicloud.com/api-eip/eip_apiBandwidth_0002.html)。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 告警ID，即告警规则的ID。

	AlarmId *string `json:"alarm_id,omitempty"`
}

func (o ListAllScalingV2PoliciesRequest) String() string {
	data, err := utils.Marshal(o)
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
	return utils.Marshal(c.value)
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
	return utils.Marshal(c.value)
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
	return utils.Marshal(c.value)
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
	return utils.Marshal(c.value)
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

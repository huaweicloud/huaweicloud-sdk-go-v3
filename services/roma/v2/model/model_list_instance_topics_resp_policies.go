package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListInstanceTopicsRespPolicies struct {
	// 是否为创建topic时所选择的应用。

	Owner *bool `json:"owner,omitempty"`
	// 集成应用key。

	UserName *string `json:"user_name,omitempty"`
	// 应用名称。

	AppName *string `json:"app_name,omitempty"`
	// 权限类型。   - all：发布+订阅   - pub：发布   - sub：订阅

	AccessPolicy *ListInstanceTopicsRespPoliciesAccessPolicy `json:"access_policy,omitempty"`
	// 权限类型对应的标签。  当权限类型是all时，发布和订阅的标签用符号“&”隔开。  当有多个标签时，标签用符号“||”隔开。

	Tag *string `json:"tag,omitempty"`
}

func (o ListInstanceTopicsRespPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTopicsRespPolicies struct{}"
	}

	return strings.Join([]string{"ListInstanceTopicsRespPolicies", string(data)}, " ")
}

type ListInstanceTopicsRespPoliciesAccessPolicy struct {
	value string
}

type ListInstanceTopicsRespPoliciesAccessPolicyEnum struct {
	ALL ListInstanceTopicsRespPoliciesAccessPolicy
	PUB ListInstanceTopicsRespPoliciesAccessPolicy
	SUB ListInstanceTopicsRespPoliciesAccessPolicy
}

func GetListInstanceTopicsRespPoliciesAccessPolicyEnum() ListInstanceTopicsRespPoliciesAccessPolicyEnum {
	return ListInstanceTopicsRespPoliciesAccessPolicyEnum{
		ALL: ListInstanceTopicsRespPoliciesAccessPolicy{
			value: "all",
		},
		PUB: ListInstanceTopicsRespPoliciesAccessPolicy{
			value: "pub",
		},
		SUB: ListInstanceTopicsRespPoliciesAccessPolicy{
			value: "sub",
		},
	}
}

func (c ListInstanceTopicsRespPoliciesAccessPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceTopicsRespPoliciesAccessPolicy) UnmarshalJSON(b []byte) error {
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

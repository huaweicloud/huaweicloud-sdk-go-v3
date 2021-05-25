package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 权限实体。
type UpdateTopicAccessPolicyReqPolicies struct {
	// 用户名称。

	UserName *string `json:"user_name,omitempty"`
	// 权限类型。 - all：拥有发布、订阅权限; - pub：拥有发布权限; - sub：拥有订阅权限。

	AccessPolicy *UpdateTopicAccessPolicyReqPoliciesAccessPolicy `json:"access_policy,omitempty"`
}

func (o UpdateTopicAccessPolicyReqPolicies) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyReqPolicies struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyReqPolicies", string(data)}, " ")
}

type UpdateTopicAccessPolicyReqPoliciesAccessPolicy struct {
	value string
}

type UpdateTopicAccessPolicyReqPoliciesAccessPolicyEnum struct {
	ALL UpdateTopicAccessPolicyReqPoliciesAccessPolicy
	PUB UpdateTopicAccessPolicyReqPoliciesAccessPolicy
	SUB UpdateTopicAccessPolicyReqPoliciesAccessPolicy
}

func GetUpdateTopicAccessPolicyReqPoliciesAccessPolicyEnum() UpdateTopicAccessPolicyReqPoliciesAccessPolicyEnum {
	return UpdateTopicAccessPolicyReqPoliciesAccessPolicyEnum{
		ALL: UpdateTopicAccessPolicyReqPoliciesAccessPolicy{
			value: "all",
		},
		PUB: UpdateTopicAccessPolicyReqPoliciesAccessPolicy{
			value: "pub",
		},
		SUB: UpdateTopicAccessPolicyReqPoliciesAccessPolicy{
			value: "sub",
		},
	}
}

func (c UpdateTopicAccessPolicyReqPoliciesAccessPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateTopicAccessPolicyReqPoliciesAccessPolicy) UnmarshalJSON(b []byte) error {
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

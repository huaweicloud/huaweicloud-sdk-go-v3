package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowTopicAccessPolicyRespPolicies struct {
	// 是否为创建topic时所选择的用户。

	Owner *bool `json:"owner,omitempty"`
	// 用户名。

	UserName *string `json:"user_name,omitempty"`
	// 权限类型。 - all：拥有发布、订阅权限; - pub：拥有发布权限; - sub：拥有订阅权限。

	AccessPolicy *ShowTopicAccessPolicyRespPoliciesAccessPolicy `json:"access_policy,omitempty"`
}

func (o ShowTopicAccessPolicyRespPolicies) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTopicAccessPolicyRespPolicies struct{}"
	}

	return strings.Join([]string{"ShowTopicAccessPolicyRespPolicies", string(data)}, " ")
}

type ShowTopicAccessPolicyRespPoliciesAccessPolicy struct {
	value string
}

type ShowTopicAccessPolicyRespPoliciesAccessPolicyEnum struct {
	ALL ShowTopicAccessPolicyRespPoliciesAccessPolicy
	PUB ShowTopicAccessPolicyRespPoliciesAccessPolicy
	SUB ShowTopicAccessPolicyRespPoliciesAccessPolicy
}

func GetShowTopicAccessPolicyRespPoliciesAccessPolicyEnum() ShowTopicAccessPolicyRespPoliciesAccessPolicyEnum {
	return ShowTopicAccessPolicyRespPoliciesAccessPolicyEnum{
		ALL: ShowTopicAccessPolicyRespPoliciesAccessPolicy{
			value: "all",
		},
		PUB: ShowTopicAccessPolicyRespPoliciesAccessPolicy{
			value: "pub",
		},
		SUB: ShowTopicAccessPolicyRespPoliciesAccessPolicy{
			value: "sub",
		},
	}
}

func (c ShowTopicAccessPolicyRespPoliciesAccessPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowTopicAccessPolicyRespPoliciesAccessPolicy) UnmarshalJSON(b []byte) error {
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

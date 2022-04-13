package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type UpdateTopicAccessPolicyPoliciesObject struct {
	// 集成应用key。

	AppId string `json:"app_id"`
	// 应用名称。

	AppName string `json:"app_name"`
	// 权限类型。   - all：发布+订阅   - pub：发布   - sub：订阅

	AccessPolicy UpdateTopicAccessPolicyPoliciesObjectAccessPolicy `json:"access_policy"`
	// 是否为创建topic时所选择的应用。  默认为false。

	Owner *bool `json:"owner,omitempty"`
	// 权限类型对应的标签。  当权限类型是all时，发布和订阅的标签用符号“&”隔开。  当有多个标签时，标签用符号“||”隔开。

	Tag *string `json:"tag,omitempty"`
}

func (o UpdateTopicAccessPolicyPoliciesObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyPoliciesObject struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyPoliciesObject", string(data)}, " ")
}

type UpdateTopicAccessPolicyPoliciesObjectAccessPolicy struct {
	value string
}

type UpdateTopicAccessPolicyPoliciesObjectAccessPolicyEnum struct {
	ALL UpdateTopicAccessPolicyPoliciesObjectAccessPolicy
	PUB UpdateTopicAccessPolicyPoliciesObjectAccessPolicy
	SUB UpdateTopicAccessPolicyPoliciesObjectAccessPolicy
}

func GetUpdateTopicAccessPolicyPoliciesObjectAccessPolicyEnum() UpdateTopicAccessPolicyPoliciesObjectAccessPolicyEnum {
	return UpdateTopicAccessPolicyPoliciesObjectAccessPolicyEnum{
		ALL: UpdateTopicAccessPolicyPoliciesObjectAccessPolicy{
			value: "all",
		},
		PUB: UpdateTopicAccessPolicyPoliciesObjectAccessPolicy{
			value: "pub",
		},
		SUB: UpdateTopicAccessPolicyPoliciesObjectAccessPolicy{
			value: "sub",
		},
	}
}

func (c UpdateTopicAccessPolicyPoliciesObjectAccessPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTopicAccessPolicyPoliciesObjectAccessPolicy) UnmarshalJSON(b []byte) error {
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

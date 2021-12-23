package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 权限实体。
type AccessPolicyEntity struct {
	// 用户名称。

	UserName *string `json:"user_name,omitempty"`
	// 权限类型。 - all：拥有发布、订阅权限; - pub：拥有发布权限; - sub：拥有订阅权限。

	AccessPolicy *AccessPolicyEntityAccessPolicy `json:"access_policy,omitempty"`
}

func (o AccessPolicyEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyEntity struct{}"
	}

	return strings.Join([]string{"AccessPolicyEntity", string(data)}, " ")
}

type AccessPolicyEntityAccessPolicy struct {
	value string
}

type AccessPolicyEntityAccessPolicyEnum struct {
	ALL AccessPolicyEntityAccessPolicy
	PUB AccessPolicyEntityAccessPolicy
	SUB AccessPolicyEntityAccessPolicy
}

func GetAccessPolicyEntityAccessPolicyEnum() AccessPolicyEntityAccessPolicyEnum {
	return AccessPolicyEntityAccessPolicyEnum{
		ALL: AccessPolicyEntityAccessPolicy{
			value: "all",
		},
		PUB: AccessPolicyEntityAccessPolicy{
			value: "pub",
		},
		SUB: AccessPolicyEntityAccessPolicy{
			value: "sub",
		},
	}
}

func (c AccessPolicyEntityAccessPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyEntityAccessPolicy) UnmarshalJSON(b []byte) error {
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

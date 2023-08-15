package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessPolicyObject struct {

	// 黑名单中的对象id。
	ObjectId *string `json:"object_id,omitempty"`

	// 对象名。
	ObjectName *string `json:"object_name,omitempty"`

	// 对象类型。 * USER： 用户 * USERGROUP： 用户组
	ObjectType *AccessPolicyObjectObjectType `json:"object_type,omitempty"`
}

func (o AccessPolicyObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyObject struct{}"
	}

	return strings.Join([]string{"AccessPolicyObject", string(data)}, " ")
}

type AccessPolicyObjectObjectType struct {
	value string
}

type AccessPolicyObjectObjectTypeEnum struct {
	USER      AccessPolicyObjectObjectType
	USERGROUP AccessPolicyObjectObjectType
}

func GetAccessPolicyObjectObjectTypeEnum() AccessPolicyObjectObjectTypeEnum {
	return AccessPolicyObjectObjectTypeEnum{
		USER: AccessPolicyObjectObjectType{
			value: "USER",
		},
		USERGROUP: AccessPolicyObjectObjectType{
			value: "USERGROUP",
		},
	}
}

func (c AccessPolicyObjectObjectType) Value() string {
	return c.value
}

func (c AccessPolicyObjectObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyObjectObjectType) UnmarshalJSON(b []byte) error {
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

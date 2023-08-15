package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessPolicyObjectInfo struct {

	// 黑名单中的对象id。
	ObjectId string `json:"object_id"`

	// 对象类型。 * USER： 用户 * USERGROUP： 用户组
	ObjectType AccessPolicyObjectInfoObjectType `json:"object_type"`

	// 对象名。后续此参数不会校验。
	ObjectName *string `json:"object_name,omitempty"`
}

func (o AccessPolicyObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyObjectInfo struct{}"
	}

	return strings.Join([]string{"AccessPolicyObjectInfo", string(data)}, " ")
}

type AccessPolicyObjectInfoObjectType struct {
	value string
}

type AccessPolicyObjectInfoObjectTypeEnum struct {
	USER      AccessPolicyObjectInfoObjectType
	USERGROUP AccessPolicyObjectInfoObjectType
}

func GetAccessPolicyObjectInfoObjectTypeEnum() AccessPolicyObjectInfoObjectTypeEnum {
	return AccessPolicyObjectInfoObjectTypeEnum{
		USER: AccessPolicyObjectInfoObjectType{
			value: "USER",
		},
		USERGROUP: AccessPolicyObjectInfoObjectType{
			value: "USERGROUP",
		},
	}
}

func (c AccessPolicyObjectInfoObjectType) Value() string {
	return c.value
}

func (c AccessPolicyObjectInfoObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyObjectInfoObjectType) UnmarshalJSON(b []byte) error {
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

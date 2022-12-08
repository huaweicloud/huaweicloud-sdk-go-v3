package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessPolicyObjectInfo struct {

	// 黑名单中的对象id，目前仅用户对象。
	ObjectId string `json:"object_id"`

	// 对象名。
	ObjectName *string `json:"object_name,omitempty"`

	// 对象类型，当前只支持用户类型。 * USER： 用户
	ObjectType AccessPolicyObjectInfoObjectType `json:"object_type"`
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
	USER AccessPolicyObjectInfoObjectType
}

func GetAccessPolicyObjectInfoObjectTypeEnum() AccessPolicyObjectInfoObjectTypeEnum {
	return AccessPolicyObjectInfoObjectTypeEnum{
		USER: AccessPolicyObjectInfoObjectType{
			value: "USER",
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

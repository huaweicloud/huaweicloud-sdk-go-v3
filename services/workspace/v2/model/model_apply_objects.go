package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplyObjects 辅助认证策略应用的用户/用户组信息。
type ApplyObjects struct {

	// 绑定对象类型枚举。 - USER：用户 - USER_GROUP：用户组 - ALL: 全部用户
	ObjectType ApplyObjectsObjectType `json:"object_type"`

	// 用户/用户组id。
	ObjectId string `json:"object_id"`

	// 用户/用户组名称。
	ObjectName string `json:"object_name"`
}

func (o ApplyObjects) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyObjects struct{}"
	}

	return strings.Join([]string{"ApplyObjects", string(data)}, " ")
}

type ApplyObjectsObjectType struct {
	value string
}

type ApplyObjectsObjectTypeEnum struct {
	USER       ApplyObjectsObjectType
	USER_GROUP ApplyObjectsObjectType
	ALL        ApplyObjectsObjectType
}

func GetApplyObjectsObjectTypeEnum() ApplyObjectsObjectTypeEnum {
	return ApplyObjectsObjectTypeEnum{
		USER: ApplyObjectsObjectType{
			value: "USER",
		},
		USER_GROUP: ApplyObjectsObjectType{
			value: "USER_GROUP",
		},
		ALL: ApplyObjectsObjectType{
			value: "ALL",
		},
	}
}

func (c ApplyObjectsObjectType) Value() string {
	return c.value
}

func (c ApplyObjectsObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplyObjectsObjectType) UnmarshalJSON(b []byte) error {
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

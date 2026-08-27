package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplyObjectInfo 应用对象信息
type ApplyObjectInfo struct {

	// 应用对象类型，包括DESKTOP（单桌面）、ALL_DESKTOPS（全部桌面）、DESKTOP_POOL（桌面池）、DESKTOP_TAG（桌面标签）、ALL_USERS（全部用户）、USER（单个用户）、USER_GROUP（用户组）
	ObjectType *ApplyObjectInfoObjectType `json:"object_type,omitempty"`

	// 对象ID（object_type为ALL_DESKTOPS或ALL_USERS时可为null）
	ObjectId *string `json:"object_id,omitempty"`
}

func (o ApplyObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyObjectInfo struct{}"
	}

	return strings.Join([]string{"ApplyObjectInfo", string(data)}, " ")
}

type ApplyObjectInfoObjectType struct {
	value string
}

type ApplyObjectInfoObjectTypeEnum struct {
	DESKTOP      ApplyObjectInfoObjectType
	ALL_DESKTOPS ApplyObjectInfoObjectType
	DESKTOP_POOL ApplyObjectInfoObjectType
	DESKTOP_TAG  ApplyObjectInfoObjectType
	ALL_USERS    ApplyObjectInfoObjectType
	USER         ApplyObjectInfoObjectType
	USER_GROUP   ApplyObjectInfoObjectType
}

func GetApplyObjectInfoObjectTypeEnum() ApplyObjectInfoObjectTypeEnum {
	return ApplyObjectInfoObjectTypeEnum{
		DESKTOP: ApplyObjectInfoObjectType{
			value: "DESKTOP",
		},
		ALL_DESKTOPS: ApplyObjectInfoObjectType{
			value: "ALL_DESKTOPS",
		},
		DESKTOP_POOL: ApplyObjectInfoObjectType{
			value: "DESKTOP_POOL",
		},
		DESKTOP_TAG: ApplyObjectInfoObjectType{
			value: "DESKTOP_TAG",
		},
		ALL_USERS: ApplyObjectInfoObjectType{
			value: "ALL_USERS",
		},
		USER: ApplyObjectInfoObjectType{
			value: "USER",
		},
		USER_GROUP: ApplyObjectInfoObjectType{
			value: "USER_GROUP",
		},
	}
}

func (c ApplyObjectInfoObjectType) Value() string {
	return c.value
}

func (c ApplyObjectInfoObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplyObjectInfoObjectType) UnmarshalJSON(b []byte) error {
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

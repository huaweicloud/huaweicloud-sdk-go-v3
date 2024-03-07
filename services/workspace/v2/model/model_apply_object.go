package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplyObject 定时任务应用对象。
type ApplyObject struct {

	// 对象ID。
	ObjectId *string `json:"object_id,omitempty"`

	// 对象类型，可选值为： - DESKTOP：桌面。 - DESKTOP_POOL：桌面池。 - ALL_DESKTOPS: 所有桌面，仅供触发式任务使用。
	ObjectType *ApplyObjectObjectType `json:"object_type,omitempty"`

	// 对象名称。
	ObjectName *string `json:"object_name,omitempty"`
}

func (o ApplyObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyObject struct{}"
	}

	return strings.Join([]string{"ApplyObject", string(data)}, " ")
}

type ApplyObjectObjectType struct {
	value string
}

type ApplyObjectObjectTypeEnum struct {
	DESKTOP      ApplyObjectObjectType
	DESKTOP_POOL ApplyObjectObjectType
	ALL_DESKTOPS ApplyObjectObjectType
}

func GetApplyObjectObjectTypeEnum() ApplyObjectObjectTypeEnum {
	return ApplyObjectObjectTypeEnum{
		DESKTOP: ApplyObjectObjectType{
			value: "DESKTOP",
		},
		DESKTOP_POOL: ApplyObjectObjectType{
			value: "DESKTOP_POOL",
		},
		ALL_DESKTOPS: ApplyObjectObjectType{
			value: "ALL_DESKTOPS",
		},
	}
}

func (c ApplyObjectObjectType) Value() string {
	return c.value
}

func (c ApplyObjectObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplyObjectObjectType) UnmarshalJSON(b []byte) error {
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

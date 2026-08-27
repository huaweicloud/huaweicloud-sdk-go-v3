package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModelConfigItem 模型分组与资源的关联项。
type ModelConfigItem struct {

	// 模型分组ID。
	GroupId string `json:"group_id"`

	// 资源ID（Agent实例ID或桌面标签key:value）。
	ResourceId string `json:"resource_id"`

	// 资源类型（DESKTOP-桌面实例，DESKTOP_TAG-桌面标签）。
	ResourceType ModelConfigItemResourceType `json:"resource_type"`
}

func (o ModelConfigItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelConfigItem struct{}"
	}

	return strings.Join([]string{"ModelConfigItem", string(data)}, " ")
}

type ModelConfigItemResourceType struct {
	value string
}

type ModelConfigItemResourceTypeEnum struct {
	DESKTOP     ModelConfigItemResourceType
	DESKTOP_TAG ModelConfigItemResourceType
}

func GetModelConfigItemResourceTypeEnum() ModelConfigItemResourceTypeEnum {
	return ModelConfigItemResourceTypeEnum{
		DESKTOP: ModelConfigItemResourceType{
			value: "DESKTOP",
		},
		DESKTOP_TAG: ModelConfigItemResourceType{
			value: "DESKTOP_TAG",
		},
	}
}

func (c ModelConfigItemResourceType) Value() string {
	return c.value
}

func (c ModelConfigItemResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModelConfigItemResourceType) UnmarshalJSON(b []byte) error {
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

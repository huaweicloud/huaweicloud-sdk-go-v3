package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateResourceLevelBody 手动改变资源等级请求体
type UpdateResourceLevelBody struct {

	// change和recover两种枚举
	ChangeType UpdateResourceLevelBodyChangeType `json:"change_type"`

	// resource_id的列表
	ResourceIds []string `json:"resource_ids"`

	// 资源等级ID
	ResourceLevelId *string `json:"resource_level_id,omitempty"`
}

func (o UpdateResourceLevelBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceLevelBody struct{}"
	}

	return strings.Join([]string{"UpdateResourceLevelBody", string(data)}, " ")
}

type UpdateResourceLevelBodyChangeType struct {
	value string
}

type UpdateResourceLevelBodyChangeTypeEnum struct {
	CHANGE  UpdateResourceLevelBodyChangeType
	RECOVER UpdateResourceLevelBodyChangeType
}

func GetUpdateResourceLevelBodyChangeTypeEnum() UpdateResourceLevelBodyChangeTypeEnum {
	return UpdateResourceLevelBodyChangeTypeEnum{
		CHANGE: UpdateResourceLevelBodyChangeType{
			value: "change",
		},
		RECOVER: UpdateResourceLevelBodyChangeType{
			value: "recover",
		},
	}
}

func (c UpdateResourceLevelBodyChangeType) Value() string {
	return c.value
}

func (c UpdateResourceLevelBodyChangeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateResourceLevelBodyChangeType) UnmarshalJSON(b []byte) error {
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

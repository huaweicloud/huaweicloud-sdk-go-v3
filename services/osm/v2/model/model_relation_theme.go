package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RelationTheme struct {

	// 主题
	ThemeName *string `json:"theme_name,omitempty"`

	// - NON_PRIMARY: 主主题 - PRIMARY:
	RelationType *RelationThemeRelationType `json:"relation_type,omitempty"`
}

func (o RelationTheme) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationTheme struct{}"
	}

	return strings.Join([]string{"RelationTheme", string(data)}, " ")
}

type RelationThemeRelationType struct {
	value string
}

type RelationThemeRelationTypeEnum struct {
	NON_PRIMARY RelationThemeRelationType
	PRIMARY     RelationThemeRelationType
}

func GetRelationThemeRelationTypeEnum() RelationThemeRelationTypeEnum {
	return RelationThemeRelationTypeEnum{
		NON_PRIMARY: RelationThemeRelationType{
			value: "NON_PRIMARY",
		},
		PRIMARY: RelationThemeRelationType{
			value: "PRIMARY",
		},
	}
}

func (c RelationThemeRelationType) Value() string {
	return c.value
}

func (c RelationThemeRelationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RelationThemeRelationType) UnmarshalJSON(b []byte) error {
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

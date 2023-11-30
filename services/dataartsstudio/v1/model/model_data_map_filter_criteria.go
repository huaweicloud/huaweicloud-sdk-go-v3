package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataMapFilterCriteria 过滤条件
type DataMapFilterCriteria struct {

	// 过滤属性维度，枚举值：typeName、base.DataAsset.sourceType、classifications.name、tags.name、securityLevel.name、workspaceId
	Attribute *DataMapFilterCriteriaAttribute `json:"attribute,omitempty"`

	// 操作表示，枚举值：EQ、IN
	Operator *DataMapFilterCriteriaOperator `json:"operator,omitempty"`

	// 属性过滤值，根据attribute变化。如attribute为数据源：base.DataAsset.sourceType，则值可为[\"dws\", \"hive\"]
	Value *[]string `json:"value,omitempty"`

	// 条件拼接准则，取值范围 AND,OR
	Condition *DataMapFilterCriteriaCondition `json:"condition,omitempty"`

	// 条件准则
	Criterion *[]DataMapFilterCriteria `json:"criterion,omitempty"`
}

func (o DataMapFilterCriteria) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataMapFilterCriteria struct{}"
	}

	return strings.Join([]string{"DataMapFilterCriteria", string(data)}, " ")
}

type DataMapFilterCriteriaAttribute struct {
	value string
}

type DataMapFilterCriteriaAttributeEnum struct {
	BASE_DATA_ASSET_SOURCE_TYPE DataMapFilterCriteriaAttribute
	TYPE_NAME                   DataMapFilterCriteriaAttribute
	CLASSIFICATIONS_NAME        DataMapFilterCriteriaAttribute
	TAGS_NAME                   DataMapFilterCriteriaAttribute
	SECURITY_LEVEL_NAME         DataMapFilterCriteriaAttribute
	WORKSPACE_ID                DataMapFilterCriteriaAttribute
}

func GetDataMapFilterCriteriaAttributeEnum() DataMapFilterCriteriaAttributeEnum {
	return DataMapFilterCriteriaAttributeEnum{
		BASE_DATA_ASSET_SOURCE_TYPE: DataMapFilterCriteriaAttribute{
			value: "base.DataAsset.sourceType",
		},
		TYPE_NAME: DataMapFilterCriteriaAttribute{
			value: "typeName",
		},
		CLASSIFICATIONS_NAME: DataMapFilterCriteriaAttribute{
			value: "classifications.name",
		},
		TAGS_NAME: DataMapFilterCriteriaAttribute{
			value: "tags.name",
		},
		SECURITY_LEVEL_NAME: DataMapFilterCriteriaAttribute{
			value: "securityLevel.name",
		},
		WORKSPACE_ID: DataMapFilterCriteriaAttribute{
			value: "workspaceId",
		},
	}
}

func (c DataMapFilterCriteriaAttribute) Value() string {
	return c.value
}

func (c DataMapFilterCriteriaAttribute) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataMapFilterCriteriaAttribute) UnmarshalJSON(b []byte) error {
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

type DataMapFilterCriteriaOperator struct {
	value string
}

type DataMapFilterCriteriaOperatorEnum struct {
	IN DataMapFilterCriteriaOperator
	EQ DataMapFilterCriteriaOperator
}

func GetDataMapFilterCriteriaOperatorEnum() DataMapFilterCriteriaOperatorEnum {
	return DataMapFilterCriteriaOperatorEnum{
		IN: DataMapFilterCriteriaOperator{
			value: "IN",
		},
		EQ: DataMapFilterCriteriaOperator{
			value: "EQ",
		},
	}
}

func (c DataMapFilterCriteriaOperator) Value() string {
	return c.value
}

func (c DataMapFilterCriteriaOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataMapFilterCriteriaOperator) UnmarshalJSON(b []byte) error {
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

type DataMapFilterCriteriaCondition struct {
	value string
}

type DataMapFilterCriteriaConditionEnum struct {
	OR DataMapFilterCriteriaCondition
}

func GetDataMapFilterCriteriaConditionEnum() DataMapFilterCriteriaConditionEnum {
	return DataMapFilterCriteriaConditionEnum{
		OR: DataMapFilterCriteriaCondition{
			value: "OR",
		},
	}
}

func (c DataMapFilterCriteriaCondition) Value() string {
	return c.value
}

func (c DataMapFilterCriteriaCondition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataMapFilterCriteriaCondition) UnmarshalJSON(b []byte) error {
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

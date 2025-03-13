package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DiagnoseTypeDto struct {

	// 数据安全诊断项 * SENSITIVE_DATA 敏感数据保护 * PERMISSION_MANAGEMENT 数据权限控制 * DATASOURCE_PROTECTION 数据源防护
	Type *DiagnoseTypeDtoType `json:"type,omitempty"`
}

func (o DiagnoseTypeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnoseTypeDto struct{}"
	}

	return strings.Join([]string{"DiagnoseTypeDto", string(data)}, " ")
}

type DiagnoseTypeDtoType struct {
	value string
}

type DiagnoseTypeDtoTypeEnum struct {
	SENSITIVE_DATA        DiagnoseTypeDtoType
	PERMISSION_MANAGEMENT DiagnoseTypeDtoType
	DATASOURCE_PROTECTION DiagnoseTypeDtoType
}

func GetDiagnoseTypeDtoTypeEnum() DiagnoseTypeDtoTypeEnum {
	return DiagnoseTypeDtoTypeEnum{
		SENSITIVE_DATA: DiagnoseTypeDtoType{
			value: "SENSITIVE_DATA",
		},
		PERMISSION_MANAGEMENT: DiagnoseTypeDtoType{
			value: "PERMISSION_MANAGEMENT",
		},
		DATASOURCE_PROTECTION: DiagnoseTypeDtoType{
			value: "DATASOURCE_PROTECTION",
		},
	}
}

func (c DiagnoseTypeDtoType) Value() string {
	return c.value
}

func (c DiagnoseTypeDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnoseTypeDtoType) UnmarshalJSON(b []byte) error {
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

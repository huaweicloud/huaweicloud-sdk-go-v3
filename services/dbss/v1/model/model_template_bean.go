package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TemplateBean 模板对象
type TemplateBean struct {

	// 数据库ID
	DbIds *string `json:"db_ids,omitempty"`

	// 数据库名称
	DbNames *string `json:"db_names,omitempty"`

	// 描述
	Desc *string `json:"desc,omitempty"`

	// 周期
	Frequency *string `json:"frequency,omitempty"`

	// 类型 - COMPREHENSIVE：综合报表 - COMPLIANCE：合规报表 - DB_SPECIAL：数据库专项报表 - CLIENT_SPECIAL：客户端专项报表 - SQL_SPECIAL：SQL
	Group *TemplateBeanGroup `json:"group,omitempty"`

	// 模板ID
	Id *string `json:"id,omitempty"`

	// 报表模板名称
	Name *string `json:"name,omitempty"`

	// 状态 - OFF：已关闭 - ON：已开启
	Status *TemplateBeanStatus `json:"status,omitempty"`

	// 报表类型 - COMPREHENSIVE：数据库安全综合报表 - COMPLIANCE：数据库安全合规报表 - SOX：SOX-萨班斯报表 - PCI：行业标准安全报表 - DB_ANALYSIS：数据库服务器分析报表 - CLIENT_IP_ANALYSIS：客户端IP分析报表 - SQL_DCL_COMMAND：DCL命令报表 - SQL_DDL_COMMAND：DDL命令报表 - SQL_DML_COMMAND：DML命令报表
	Type *TemplateBeanType `json:"type,omitempty"`
}

func (o TemplateBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateBean struct{}"
	}

	return strings.Join([]string{"TemplateBean", string(data)}, " ")
}

type TemplateBeanGroup struct {
	value string
}

type TemplateBeanGroupEnum struct {
	COMPREHENSIVE  TemplateBeanGroup
	COMPLIANCE     TemplateBeanGroup
	DB_SPECIAL     TemplateBeanGroup
	CLIENT_SPECIAL TemplateBeanGroup
	SQL_SPECIAL    TemplateBeanGroup
}

func GetTemplateBeanGroupEnum() TemplateBeanGroupEnum {
	return TemplateBeanGroupEnum{
		COMPREHENSIVE: TemplateBeanGroup{
			value: "COMPREHENSIVE",
		},
		COMPLIANCE: TemplateBeanGroup{
			value: "COMPLIANCE",
		},
		DB_SPECIAL: TemplateBeanGroup{
			value: "DB_SPECIAL",
		},
		CLIENT_SPECIAL: TemplateBeanGroup{
			value: "CLIENT_SPECIAL",
		},
		SQL_SPECIAL: TemplateBeanGroup{
			value: "SQL_SPECIAL",
		},
	}
}

func (c TemplateBeanGroup) Value() string {
	return c.value
}

func (c TemplateBeanGroup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateBeanGroup) UnmarshalJSON(b []byte) error {
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

type TemplateBeanStatus struct {
	value string
}

type TemplateBeanStatusEnum struct {
	OFF TemplateBeanStatus
	ON  TemplateBeanStatus
}

func GetTemplateBeanStatusEnum() TemplateBeanStatusEnum {
	return TemplateBeanStatusEnum{
		OFF: TemplateBeanStatus{
			value: "OFF",
		},
		ON: TemplateBeanStatus{
			value: "ON",
		},
	}
}

func (c TemplateBeanStatus) Value() string {
	return c.value
}

func (c TemplateBeanStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateBeanStatus) UnmarshalJSON(b []byte) error {
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

type TemplateBeanType struct {
	value string
}

type TemplateBeanTypeEnum struct {
	COMPREHENSIVE      TemplateBeanType
	COMPLIANCE         TemplateBeanType
	SOX                TemplateBeanType
	PCI                TemplateBeanType
	DB_ANALYSIS        TemplateBeanType
	CLIENT_IP_ANALYSIS TemplateBeanType
	SQL_DCL_COMMAND    TemplateBeanType
	SQL_DDL_COMMAND    TemplateBeanType
	SQL_DML_COMMAND    TemplateBeanType
}

func GetTemplateBeanTypeEnum() TemplateBeanTypeEnum {
	return TemplateBeanTypeEnum{
		COMPREHENSIVE: TemplateBeanType{
			value: "COMPREHENSIVE",
		},
		COMPLIANCE: TemplateBeanType{
			value: "COMPLIANCE",
		},
		SOX: TemplateBeanType{
			value: "SOX",
		},
		PCI: TemplateBeanType{
			value: "PCI",
		},
		DB_ANALYSIS: TemplateBeanType{
			value: "DB_ANALYSIS",
		},
		CLIENT_IP_ANALYSIS: TemplateBeanType{
			value: "CLIENT_IP_ANALYSIS",
		},
		SQL_DCL_COMMAND: TemplateBeanType{
			value: "SQL_DCL_COMMAND",
		},
		SQL_DDL_COMMAND: TemplateBeanType{
			value: "SQL_DDL_COMMAND",
		},
		SQL_DML_COMMAND: TemplateBeanType{
			value: "SQL_DML_COMMAND",
		},
	}
}

func (c TemplateBeanType) Value() string {
	return c.value
}

func (c TemplateBeanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateBeanType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BizTypeEnum 业务实体类型。AGGREGATION_LOGIC_TABL(汇总表)、ATOMIC_INDE(原子指标)、ATOMIC_METRI(原子指标（新）)、BIZ_CATALO(流程架构目录)、BIZ_METRI(业务指标)、CODE_TABL(码表)、COMMON_CONDITIO(通用限定)、COMPOSITE_METRI(复合指标（新）)、COMPOUND_METRI(复合指标)、CONDITION_GROU(限定分组)、DEGENERATE_DIMENSIO(退化维度)、DERIVATIVE_INDE(衍生指标)、DERIVED_METRI(衍生指标（新）)、DIMENSIO(维度)、DIMENSION_ATTRIBUT(维度属性)、DIMENSION_HIERARCHIE(维度层级)、DIMENSION_LOGIC_TABL(维度表)、DIMENSION_TABLE_ATTRIBUT(维度属性)、DIRECTOR(目录)、FACT_ATTRIBUT(事实表属性)、FACT_DIMENSIO(事实表维度)、FACT_LOGIC_TABL(事实表)、FACT_MEASUR(事实表度量)、FUNCTIO(函数)、INFO_ARC(信息架构（批量修改主题使用）)、MODE(模型)、QUALITY_RUL(质量规则)、SECRECY_LEVE(密级)、STANDARD_ELEMEN(数据标准)、STANDARD_ELEMENT_TEMPLAT(数据标准模板)、SUBJEC(主题)、SUMMARY_DIMENSION_ATTRIBUT(汇总表维度属性)、SUMMARY_INDE(汇总表指标属性)、SUMMARY_TIM(汇总表时间周期属性)、TABLE_MODE(关系模型（逻辑模型/物理模型）)、TABLE_MODEL_ATTRIBUT(关系模型属性（逻辑模型/物理模型）)、TABLE_MODEL_LOGI(逻辑实体)、TABLE_TYP(表类型)、TA(标签)、TIME_CONDITIO(时间限定)。
type BizTypeEnum struct {
	value string
}

type BizTypeEnumEnum struct {
	AGGREGATION_LOGIC_TABLE     BizTypeEnum
	ATOMIC_INDEX                BizTypeEnum
	ATOMIC_METRIC               BizTypeEnum
	BIZ_CATALOG                 BizTypeEnum
	BIZ_METRIC                  BizTypeEnum
	CODE_TABLE                  BizTypeEnum
	COMMON_CONDITION            BizTypeEnum
	COMPOSITE_METRIC            BizTypeEnum
	COMPOUND_METRIC             BizTypeEnum
	CONDITION_GROUP             BizTypeEnum
	DEGENERATE_DIMENSION        BizTypeEnum
	DERIVATIVE_INDEX            BizTypeEnum
	DERIVED_METRIC              BizTypeEnum
	DIMENSION                   BizTypeEnum
	DIMENSION_ATTRIBUTE         BizTypeEnum
	DIMENSION_HIERARCHIES       BizTypeEnum
	DIMENSION_LOGIC_TABLE       BizTypeEnum
	DIMENSION_TABLE_ATTRIBUTE   BizTypeEnum
	DIRECTORY                   BizTypeEnum
	FACT_ATTRIBUTE              BizTypeEnum
	FACT_DIMENSION              BizTypeEnum
	FACT_LOGIC_TABLE            BizTypeEnum
	FACT_MEASURE                BizTypeEnum
	FUNCTION                    BizTypeEnum
	INFO_ARCH                   BizTypeEnum
	MODEL                       BizTypeEnum
	QUALITY_RULE                BizTypeEnum
	SECRECY_LEVEL               BizTypeEnum
	STANDARD_ELEMENT            BizTypeEnum
	STANDARD_ELEMENT_TEMPLATE   BizTypeEnum
	SUBJECT                     BizTypeEnum
	SUMMARY_DIMENSION_ATTRIBUTE BizTypeEnum
	SUMMARY_INDEX               BizTypeEnum
	SUMMARY_TIME                BizTypeEnum
	TABLE_MODEL                 BizTypeEnum
	TABLE_MODEL_ATTRIBUTE       BizTypeEnum
	TABLE_MODEL_LOGIC           BizTypeEnum
	TABLE_TYPE                  BizTypeEnum
	TAG                         BizTypeEnum
	TIME_CONDITION              BizTypeEnum
}

func GetBizTypeEnumEnum() BizTypeEnumEnum {
	return BizTypeEnumEnum{
		AGGREGATION_LOGIC_TABLE: BizTypeEnum{
			value: "AGGREGATION_LOGIC_TABLE",
		},
		ATOMIC_INDEX: BizTypeEnum{
			value: "ATOMIC_INDEX",
		},
		ATOMIC_METRIC: BizTypeEnum{
			value: "ATOMIC_METRIC",
		},
		BIZ_CATALOG: BizTypeEnum{
			value: "BIZ_CATALOG",
		},
		BIZ_METRIC: BizTypeEnum{
			value: "BIZ_METRIC",
		},
		CODE_TABLE: BizTypeEnum{
			value: "CODE_TABLE",
		},
		COMMON_CONDITION: BizTypeEnum{
			value: "COMMON_CONDITION",
		},
		COMPOSITE_METRIC: BizTypeEnum{
			value: "COMPOSITE_METRIC",
		},
		COMPOUND_METRIC: BizTypeEnum{
			value: "COMPOUND_METRIC",
		},
		CONDITION_GROUP: BizTypeEnum{
			value: "CONDITION_GROUP",
		},
		DEGENERATE_DIMENSION: BizTypeEnum{
			value: "DEGENERATE_DIMENSION",
		},
		DERIVATIVE_INDEX: BizTypeEnum{
			value: "DERIVATIVE_INDEX",
		},
		DERIVED_METRIC: BizTypeEnum{
			value: "DERIVED_METRIC",
		},
		DIMENSION: BizTypeEnum{
			value: "DIMENSION",
		},
		DIMENSION_ATTRIBUTE: BizTypeEnum{
			value: "DIMENSION_ATTRIBUTE",
		},
		DIMENSION_HIERARCHIES: BizTypeEnum{
			value: "DIMENSION_HIERARCHIES",
		},
		DIMENSION_LOGIC_TABLE: BizTypeEnum{
			value: "DIMENSION_LOGIC_TABLE",
		},
		DIMENSION_TABLE_ATTRIBUTE: BizTypeEnum{
			value: "DIMENSION_TABLE_ATTRIBUTE",
		},
		DIRECTORY: BizTypeEnum{
			value: "DIRECTORY",
		},
		FACT_ATTRIBUTE: BizTypeEnum{
			value: "FACT_ATTRIBUTE",
		},
		FACT_DIMENSION: BizTypeEnum{
			value: "FACT_DIMENSION",
		},
		FACT_LOGIC_TABLE: BizTypeEnum{
			value: "FACT_LOGIC_TABLE",
		},
		FACT_MEASURE: BizTypeEnum{
			value: "FACT_MEASURE",
		},
		FUNCTION: BizTypeEnum{
			value: "FUNCTION",
		},
		INFO_ARCH: BizTypeEnum{
			value: "INFO_ARCH",
		},
		MODEL: BizTypeEnum{
			value: "MODEL",
		},
		QUALITY_RULE: BizTypeEnum{
			value: "QUALITY_RULE",
		},
		SECRECY_LEVEL: BizTypeEnum{
			value: "SECRECY_LEVEL",
		},
		STANDARD_ELEMENT: BizTypeEnum{
			value: "STANDARD_ELEMENT",
		},
		STANDARD_ELEMENT_TEMPLATE: BizTypeEnum{
			value: "STANDARD_ELEMENT_TEMPLATE",
		},
		SUBJECT: BizTypeEnum{
			value: "SUBJECT",
		},
		SUMMARY_DIMENSION_ATTRIBUTE: BizTypeEnum{
			value: "SUMMARY_DIMENSION_ATTRIBUTE",
		},
		SUMMARY_INDEX: BizTypeEnum{
			value: "SUMMARY_INDEX",
		},
		SUMMARY_TIME: BizTypeEnum{
			value: "SUMMARY_TIME",
		},
		TABLE_MODEL: BizTypeEnum{
			value: "TABLE_MODEL",
		},
		TABLE_MODEL_ATTRIBUTE: BizTypeEnum{
			value: "TABLE_MODEL_ATTRIBUTE",
		},
		TABLE_MODEL_LOGIC: BizTypeEnum{
			value: "TABLE_MODEL_LOGIC",
		},
		TABLE_TYPE: BizTypeEnum{
			value: "TABLE_TYPE",
		},
		TAG: BizTypeEnum{
			value: "TAG",
		},
		TIME_CONDITION: BizTypeEnum{
			value: "TIME_CONDITION",
		},
	}
}

func (c BizTypeEnum) Value() string {
	return c.value
}

func (c BizTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BizTypeEnum) UnmarshalJSON(b []byte) error {
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

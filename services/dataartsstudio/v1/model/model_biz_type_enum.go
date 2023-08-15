package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BizTypeEnum 业务类型
type BizTypeEnum struct {
	value string
}

type BizTypeEnumEnum struct {
	ATOMIC_INDEX                BizTypeEnum
	DERIVATIVE_INDEX            BizTypeEnum
	DIMENSION                   BizTypeEnum
	DIMENSION_ATTRIBUTE         BizTypeEnum
	DIMENSION_HIERARCHIES       BizTypeEnum
	CONDITION_GROUP             BizTypeEnum
	TIME_CONDITION              BizTypeEnum
	COMMON_CONDITION            BizTypeEnum
	FACT_LOGIC_TABLE            BizTypeEnum
	TABLE_MODEL                 BizTypeEnum
	DIMENSION_LOGIC_TABLE       BizTypeEnum
	STANDARD_ELEMENT            BizTypeEnum
	AGGREGATION_LOGIC_TABLE     BizTypeEnum
	TAG                         BizTypeEnum
	CODE_TABLE                  BizTypeEnum
	TABLE_MODEL_ATTRIBUTE       BizTypeEnum
	FACT_DIMENSION              BizTypeEnum
	FACT_ATTRIBUTE              BizTypeEnum
	FACT_MEASURE                BizTypeEnum
	SUMMARY_DIMENSION_ATTRIBUTE BizTypeEnum
	SUMMARY_TIME                BizTypeEnum
	DIMENSION_TABLE_ATTRIBUTE   BizTypeEnum
	QUALITY_RULE                BizTypeEnum
	BIZ_METRIC                  BizTypeEnum
	COMPOUND_METRIC             BizTypeEnum
	INFO_ARCH                   BizTypeEnum
	DEGENERATE_DIMENSION        BizTypeEnum
	TABLE_MODEL_LOGIC           BizTypeEnum
}

func GetBizTypeEnumEnum() BizTypeEnumEnum {
	return BizTypeEnumEnum{
		ATOMIC_INDEX: BizTypeEnum{
			value: "ATOMIC_INDEX",
		},
		DERIVATIVE_INDEX: BizTypeEnum{
			value: "DERIVATIVE_INDEX",
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
		CONDITION_GROUP: BizTypeEnum{
			value: "CONDITION_GROUP",
		},
		TIME_CONDITION: BizTypeEnum{
			value: "TIME_CONDITION",
		},
		COMMON_CONDITION: BizTypeEnum{
			value: "COMMON_CONDITION",
		},
		FACT_LOGIC_TABLE: BizTypeEnum{
			value: "FACT_LOGIC_TABLE",
		},
		TABLE_MODEL: BizTypeEnum{
			value: "TABLE_MODEL",
		},
		DIMENSION_LOGIC_TABLE: BizTypeEnum{
			value: "DIMENSION_LOGIC_TABLE",
		},
		STANDARD_ELEMENT: BizTypeEnum{
			value: "STANDARD_ELEMENT",
		},
		AGGREGATION_LOGIC_TABLE: BizTypeEnum{
			value: "AGGREGATION_LOGIC_TABLE",
		},
		TAG: BizTypeEnum{
			value: "TAG",
		},
		CODE_TABLE: BizTypeEnum{
			value: "CODE_TABLE",
		},
		TABLE_MODEL_ATTRIBUTE: BizTypeEnum{
			value: "TABLE_MODEL_ATTRIBUTE",
		},
		FACT_DIMENSION: BizTypeEnum{
			value: "FACT_DIMENSION",
		},
		FACT_ATTRIBUTE: BizTypeEnum{
			value: "FACT_ATTRIBUTE",
		},
		FACT_MEASURE: BizTypeEnum{
			value: "FACT_MEASURE",
		},
		SUMMARY_DIMENSION_ATTRIBUTE: BizTypeEnum{
			value: "SUMMARY_DIMENSION_ATTRIBUTE",
		},
		SUMMARY_TIME: BizTypeEnum{
			value: "SUMMARY_TIME",
		},
		DIMENSION_TABLE_ATTRIBUTE: BizTypeEnum{
			value: "DIMENSION_TABLE_ATTRIBUTE",
		},
		QUALITY_RULE: BizTypeEnum{
			value: "QUALITY_RULE",
		},
		BIZ_METRIC: BizTypeEnum{
			value: "BIZ_METRIC",
		},
		COMPOUND_METRIC: BizTypeEnum{
			value: "COMPOUND_METRIC",
		},
		INFO_ARCH: BizTypeEnum{
			value: "INFO_ARCH",
		},
		DEGENERATE_DIMENSION: BizTypeEnum{
			value: "DEGENERATE_DIMENSION",
		},
		TABLE_MODEL_LOGIC: BizTypeEnum{
			value: "TABLE_MODEL_LOGIC",
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

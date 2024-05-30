package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateDesignTableQualityParam struct {

	// 表类型。 枚举值：   - TABLE_MODEL: 关系模型（逻辑模型/物理模型）   - AGGREGATION_LOGIC_TABLE: 汇总表   - FACT_LOGIC_TABLE: 事实表   - DIMENSION: 维度   - DIMENSION_LOGIC_TABLE: 维度表
	BizType *UpdateDesignTableQualityParamBizType `json:"biz_type,omitempty"`

	// 异常数据输出开关。
	DirtyOutSwitch *bool `json:"dirty_out_switch,omitempty"`

	// 异常数据输出库。
	DirtyOutDatabase *string `json:"dirty_out_database,omitempty"`

	// 异常表前缀。
	DirtyOutPrefix *string `json:"dirty_out_prefix,omitempty"`

	// 异常表后缀。
	DirtyOutSuffix *string `json:"dirty_out_suffix,omitempty"`
}

func (o UpdateDesignTableQualityParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignTableQualityParam struct{}"
	}

	return strings.Join([]string{"UpdateDesignTableQualityParam", string(data)}, " ")
}

type UpdateDesignTableQualityParamBizType struct {
	value string
}

type UpdateDesignTableQualityParamBizTypeEnum struct {
	TABLE_MODEL             UpdateDesignTableQualityParamBizType
	AGGREGATION_LOGIC_TABLE UpdateDesignTableQualityParamBizType
	FACT_LOGIC_TABLE        UpdateDesignTableQualityParamBizType
	DIMENSION               UpdateDesignTableQualityParamBizType
	DIMENSION_LOGIC_TABLE   UpdateDesignTableQualityParamBizType
}

func GetUpdateDesignTableQualityParamBizTypeEnum() UpdateDesignTableQualityParamBizTypeEnum {
	return UpdateDesignTableQualityParamBizTypeEnum{
		TABLE_MODEL: UpdateDesignTableQualityParamBizType{
			value: "TABLE_MODEL",
		},
		AGGREGATION_LOGIC_TABLE: UpdateDesignTableQualityParamBizType{
			value: "AGGREGATION_LOGIC_TABLE",
		},
		FACT_LOGIC_TABLE: UpdateDesignTableQualityParamBizType{
			value: "FACT_LOGIC_TABLE",
		},
		DIMENSION: UpdateDesignTableQualityParamBizType{
			value: "DIMENSION",
		},
		DIMENSION_LOGIC_TABLE: UpdateDesignTableQualityParamBizType{
			value: "DIMENSION_LOGIC_TABLE",
		},
	}
}

func (c UpdateDesignTableQualityParamBizType) Value() string {
	return c.value
}

func (c UpdateDesignTableQualityParamBizType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDesignTableQualityParamBizType) UnmarshalJSON(b []byte) error {
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

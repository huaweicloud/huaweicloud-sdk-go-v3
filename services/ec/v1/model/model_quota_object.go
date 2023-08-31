package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QuotaObject struct {

	// 配额类型
	Type *QuotaObjectType `json:"type,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 最小值
	Min *int32 `json:"min,omitempty"`

	// 最大值
	Max *int32 `json:"max,omitempty"`

	// 配额上限
	Quota *int32 `json:"quota,omitempty"`

	// 已使用数
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaObject struct{}"
	}

	return strings.Join([]string{"QuotaObject", string(data)}, " ")
}

type QuotaObjectType struct {
	value string
}

type QuotaObjectTypeEnum struct {
	IEG_NUM_LIMIT                    QuotaObjectType
	ECN_NUM_LIMIT                    QuotaObjectType
	ECN_STANDARD_BANDWIDTH_LIMIT     QuotaObjectType
	ECN_PROFESSIONAL_BANDWIDTH_LIMIT QuotaObjectType
	ROUTE_LIMIT                      QuotaObjectType
}

func GetQuotaObjectTypeEnum() QuotaObjectTypeEnum {
	return QuotaObjectTypeEnum{
		IEG_NUM_LIMIT: QuotaObjectType{
			value: "ieg_num_limit",
		},
		ECN_NUM_LIMIT: QuotaObjectType{
			value: "ecn_num_limit",
		},
		ECN_STANDARD_BANDWIDTH_LIMIT: QuotaObjectType{
			value: "ecn_standard_bandwidth_limit",
		},
		ECN_PROFESSIONAL_BANDWIDTH_LIMIT: QuotaObjectType{
			value: "ecn_professional_bandwidth_limit",
		},
		ROUTE_LIMIT: QuotaObjectType{
			value: "route_limit",
		},
	}
}

func (c QuotaObjectType) Value() string {
	return c.value
}

func (c QuotaObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuotaObjectType) UnmarshalJSON(b []byte) error {
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

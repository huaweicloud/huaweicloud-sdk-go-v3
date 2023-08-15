package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SupplyOption 供给推荐的选项
type SupplyOption struct {

	// 推荐结果的粒度 BY_REGION：对每个区域打分，可使用多种规格满足需求 BY_AZ：对每个可用区打分 BY_FLAVOR：对每个规格打分，可使用多地域满足需求 BY_FLAVOR_AND_REGION：对每个区域下的每个规格打分 BY_FLAVOR_AND_AZ：对每个可用区下的每个规格打分
	ResultGranularity *SupplyOptionResultGranularity `json:"result_granularity,omitempty"`

	// 是否推荐竞价实例
	EnableSpot *bool `json:"enable_spot,omitempty"`
}

func (o SupplyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupplyOption struct{}"
	}

	return strings.Join([]string{"SupplyOption", string(data)}, " ")
}

type SupplyOptionResultGranularity struct {
	value string
}

type SupplyOptionResultGranularityEnum struct {
	BY_REGION            SupplyOptionResultGranularity
	BY_AZ                SupplyOptionResultGranularity
	BY_FLAVOR            SupplyOptionResultGranularity
	BY_FLAVOR_AND_REGION SupplyOptionResultGranularity
	BY_FLAVOR_AND_AZ     SupplyOptionResultGranularity
}

func GetSupplyOptionResultGranularityEnum() SupplyOptionResultGranularityEnum {
	return SupplyOptionResultGranularityEnum{
		BY_REGION: SupplyOptionResultGranularity{
			value: "BY_REGION",
		},
		BY_AZ: SupplyOptionResultGranularity{
			value: "BY_AZ",
		},
		BY_FLAVOR: SupplyOptionResultGranularity{
			value: "BY_FLAVOR",
		},
		BY_FLAVOR_AND_REGION: SupplyOptionResultGranularity{
			value: "BY_FLAVOR_AND_REGION",
		},
		BY_FLAVOR_AND_AZ: SupplyOptionResultGranularity{
			value: "BY_FLAVOR_AND_AZ",
		},
	}
}

func (c SupplyOptionResultGranularity) Value() string {
	return c.value
}

func (c SupplyOptionResultGranularity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SupplyOptionResultGranularity) UnmarshalJSON(b []byte) error {
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

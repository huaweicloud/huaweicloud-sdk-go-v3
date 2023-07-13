package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupplyRecommendationRequestBody This is a auto create Body Object
type ListSupplyRecommendationRequestBody struct {
	FlavorConstraint *FlavorConstraint `json:"flavor_constraint,omitempty"`

	// 接受推荐的规格列表
	FlavorIds *[]string `json:"flavor_ids,omitempty"`

	// 接受推荐的地域列表，默认接受所有区域
	Locations *[]DistinctLocation `json:"locations,omitempty"`

	Option *SupplyOption `json:"option,omitempty"`

	// 推荐的策略。 CAPACITY：容量策略 COST：成本策略
	Strategy *ListSupplyRecommendationRequestBodyStrategy `json:"strategy,omitempty"`

	// 查询返回的数量限制
	Limit *int32 `json:"limit,omitempty"`

	// 取值为上一页数据的最后一条记录的唯一标记
	Marker *string `json:"marker,omitempty"`
}

func (o ListSupplyRecommendationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupplyRecommendationRequestBody struct{}"
	}

	return strings.Join([]string{"ListSupplyRecommendationRequestBody", string(data)}, " ")
}

type ListSupplyRecommendationRequestBodyStrategy struct {
	value string
}

type ListSupplyRecommendationRequestBodyStrategyEnum struct {
	CAPACITY ListSupplyRecommendationRequestBodyStrategy
	COST     ListSupplyRecommendationRequestBodyStrategy
}

func GetListSupplyRecommendationRequestBodyStrategyEnum() ListSupplyRecommendationRequestBodyStrategyEnum {
	return ListSupplyRecommendationRequestBodyStrategyEnum{
		CAPACITY: ListSupplyRecommendationRequestBodyStrategy{
			value: "CAPACITY",
		},
		COST: ListSupplyRecommendationRequestBodyStrategy{
			value: "COST",
		},
	}
}

func (c ListSupplyRecommendationRequestBodyStrategy) Value() string {
	return c.value
}

func (c ListSupplyRecommendationRequestBodyStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupplyRecommendationRequestBodyStrategy) UnmarshalJSON(b []byte) error {
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

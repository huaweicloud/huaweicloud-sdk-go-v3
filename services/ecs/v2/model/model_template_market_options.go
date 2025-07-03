package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TemplateMarketOptions 计费条件
type TemplateMarketOptions struct {

	// 计费类型
	MarketType *TemplateMarketOptionsMarketType `json:"market_type,omitempty"`

	SpotOptions *TemplateSpotOptions `json:"spot_options,omitempty"`
}

func (o TemplateMarketOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateMarketOptions struct{}"
	}

	return strings.Join([]string{"TemplateMarketOptions", string(data)}, " ")
}

type TemplateMarketOptionsMarketType struct {
	value string
}

type TemplateMarketOptionsMarketTypeEnum struct {
	SPOT     TemplateMarketOptionsMarketType
	POSTPAID TemplateMarketOptionsMarketType
}

func GetTemplateMarketOptionsMarketTypeEnum() TemplateMarketOptionsMarketTypeEnum {
	return TemplateMarketOptionsMarketTypeEnum{
		SPOT: TemplateMarketOptionsMarketType{
			value: "spot",
		},
		POSTPAID: TemplateMarketOptionsMarketType{
			value: "postpaid",
		},
	}
}

func (c TemplateMarketOptionsMarketType) Value() string {
	return c.value
}

func (c TemplateMarketOptionsMarketType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateMarketOptionsMarketType) UnmarshalJSON(b []byte) error {
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

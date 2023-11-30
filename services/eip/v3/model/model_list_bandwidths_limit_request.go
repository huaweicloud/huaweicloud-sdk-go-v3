package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBandwidthsLimitRequest Request Object
type ListBandwidthsLimitRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 只显示指定的字段。使用ext-fields时在默认显示的字段基础上追加字段
	Fields *[]ListBandwidthsLimitRequestFields `json:"fields,omitempty"`

	// 根据charge_mode过滤
	ChargeMode *string `json:"charge_mode,omitempty"`
}

func (o ListBandwidthsLimitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthsLimitRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthsLimitRequest", string(data)}, " ")
}

type ListBandwidthsLimitRequestFields struct {
	value string
}

type ListBandwidthsLimitRequestFieldsEnum struct {
	ID          ListBandwidthsLimitRequestFields
	CHARGE_MODE ListBandwidthsLimitRequestFields
	MIN_SIZE    ListBandwidthsLimitRequestFields
	MAX_SIZE    ListBandwidthsLimitRequestFields
	EXT_LIMIT   ListBandwidthsLimitRequestFields
}

func GetListBandwidthsLimitRequestFieldsEnum() ListBandwidthsLimitRequestFieldsEnum {
	return ListBandwidthsLimitRequestFieldsEnum{
		ID: ListBandwidthsLimitRequestFields{
			value: "id",
		},
		CHARGE_MODE: ListBandwidthsLimitRequestFields{
			value: "charge_mode",
		},
		MIN_SIZE: ListBandwidthsLimitRequestFields{
			value: "min_size",
		},
		MAX_SIZE: ListBandwidthsLimitRequestFields{
			value: "max_size",
		},
		EXT_LIMIT: ListBandwidthsLimitRequestFields{
			value: "ext_limit",
		},
	}
}

func (c ListBandwidthsLimitRequestFields) Value() string {
	return c.value
}

func (c ListBandwidthsLimitRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBandwidthsLimitRequestFields) UnmarshalJSON(b []byte) error {
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

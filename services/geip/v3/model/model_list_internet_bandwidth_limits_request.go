package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInternetBandwidthLimitsRequest Request Object
type ListInternetBandwidthLimitsRequest struct {
	Fields *[]ListInternetBandwidthLimitsRequestFields `json:"fields,omitempty"`

	// 按照sort_key指定的字段排序
	SortKey *[]ListInternetBandwidthLimitsRequestSortKey `json:"sort_key,omitempty"`

	// 排序的方向，倒序或者正序
	SortDir *[]ListInternetBandwidthLimitsRequestSortDir `json:"sort_dir,omitempty"`

	ChargeMode *[]ListInternetBandwidthLimitsRequestChargeMode `json:"charge_mode,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o ListInternetBandwidthLimitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthLimitsRequest struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthLimitsRequest", string(data)}, " ")
}

type ListInternetBandwidthLimitsRequestFields struct {
	value string
}

type ListInternetBandwidthLimitsRequestFieldsEnum struct {
	ID          ListInternetBandwidthLimitsRequestFields
	CHARGE_MODE ListInternetBandwidthLimitsRequestFields
	MIN_SIZE    ListInternetBandwidthLimitsRequestFields
	EXT_LIMIT   ListInternetBandwidthLimitsRequestFields
	MAX_SIZE    ListInternetBandwidthLimitsRequestFields
}

func GetListInternetBandwidthLimitsRequestFieldsEnum() ListInternetBandwidthLimitsRequestFieldsEnum {
	return ListInternetBandwidthLimitsRequestFieldsEnum{
		ID: ListInternetBandwidthLimitsRequestFields{
			value: "id",
		},
		CHARGE_MODE: ListInternetBandwidthLimitsRequestFields{
			value: "charge_mode",
		},
		MIN_SIZE: ListInternetBandwidthLimitsRequestFields{
			value: "min_size",
		},
		EXT_LIMIT: ListInternetBandwidthLimitsRequestFields{
			value: "ext_limit",
		},
		MAX_SIZE: ListInternetBandwidthLimitsRequestFields{
			value: "max_size",
		},
	}
}

func (c ListInternetBandwidthLimitsRequestFields) Value() string {
	return c.value
}

func (c ListInternetBandwidthLimitsRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInternetBandwidthLimitsRequestFields) UnmarshalJSON(b []byte) error {
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

type ListInternetBandwidthLimitsRequestSortKey struct {
	value string
}

type ListInternetBandwidthLimitsRequestSortKeyEnum struct {
	ID ListInternetBandwidthLimitsRequestSortKey
}

func GetListInternetBandwidthLimitsRequestSortKeyEnum() ListInternetBandwidthLimitsRequestSortKeyEnum {
	return ListInternetBandwidthLimitsRequestSortKeyEnum{
		ID: ListInternetBandwidthLimitsRequestSortKey{
			value: "id",
		},
	}
}

func (c ListInternetBandwidthLimitsRequestSortKey) Value() string {
	return c.value
}

func (c ListInternetBandwidthLimitsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInternetBandwidthLimitsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListInternetBandwidthLimitsRequestSortDir struct {
	value string
}

type ListInternetBandwidthLimitsRequestSortDirEnum struct {
	ASC  ListInternetBandwidthLimitsRequestSortDir
	DESC ListInternetBandwidthLimitsRequestSortDir
}

func GetListInternetBandwidthLimitsRequestSortDirEnum() ListInternetBandwidthLimitsRequestSortDirEnum {
	return ListInternetBandwidthLimitsRequestSortDirEnum{
		ASC: ListInternetBandwidthLimitsRequestSortDir{
			value: "asc",
		},
		DESC: ListInternetBandwidthLimitsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListInternetBandwidthLimitsRequestSortDir) Value() string {
	return c.value
}

func (c ListInternetBandwidthLimitsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInternetBandwidthLimitsRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListInternetBandwidthLimitsRequestChargeMode struct {
	value string
}

type ListInternetBandwidthLimitsRequestChargeModeEnum struct {
	BANDWIDTH            ListInternetBandwidthLimitsRequestChargeMode
	E_95PEAK_BIDIRECTION ListInternetBandwidthLimitsRequestChargeMode
	E_95PEAK_PLUS_1000   ListInternetBandwidthLimitsRequestChargeMode
	E_95PEAK_GUAR        ListInternetBandwidthLimitsRequestChargeMode
}

func GetListInternetBandwidthLimitsRequestChargeModeEnum() ListInternetBandwidthLimitsRequestChargeModeEnum {
	return ListInternetBandwidthLimitsRequestChargeModeEnum{
		BANDWIDTH: ListInternetBandwidthLimitsRequestChargeMode{
			value: "bandwidth",
		},
		E_95PEAK_BIDIRECTION: ListInternetBandwidthLimitsRequestChargeMode{
			value: "95peak_bidirection",
		},
		E_95PEAK_PLUS_1000: ListInternetBandwidthLimitsRequestChargeMode{
			value: "95peak_plus_1000",
		},
		E_95PEAK_GUAR: ListInternetBandwidthLimitsRequestChargeMode{
			value: "95peak_guar",
		},
	}
}

func (c ListInternetBandwidthLimitsRequestChargeMode) Value() string {
	return c.value
}

func (c ListInternetBandwidthLimitsRequestChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInternetBandwidthLimitsRequestChargeMode) UnmarshalJSON(b []byte) error {
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

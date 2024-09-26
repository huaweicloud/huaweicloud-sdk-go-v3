package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGlobalConnectionBandwidthConfigs 租户购买全域互联带宽的动态配置项。
type ListGlobalConnectionBandwidthConfigs struct {

	// 计费类型对应全域互联带宽大小范围。
	SizeRange []GlobalConnectionBandwidthSizeRange `json:"size_range"`

	// 支持的计费类型列表。
	ChargeMode []ListGlobalConnectionBandwidthConfigsChargeMode `json:"charge_mode"`

	// 支持服务实例类型。
	Services []ListGlobalConnectionBandwidthConfigsServices `json:"services"`

	// 支持的带宽类型。
	GcbType []ListGlobalConnectionBandwidthConfigsGcbType `json:"gcb_type"`

	// 按增强型95计费保底消费百分比。
	Ratio95peakPlus *int32 `json:"ratio_95peak_plus,omitempty"`

	// 按传统型95计费保底消费百分比。
	Ratio95peakGuar *int32 `json:"ratio_95peak_guar,omitempty"`

	// 是否已经完成跨境审批。
	Crossborder bool `json:"crossborder"`

	// 配额信息。
	Quotas []GlobalConnectionBandwidthQuotas `json:"quotas"`

	// 支持线路分级。
	SlaLevel []ListGlobalConnectionBandwidthConfigsSlaLevel `json:"sla_level"`

	// 共享带宽允许绑定实例数量上限。
	BindLimit int32 `json:"bind_limit"`

	// 是否启用传统的大区带宽。
	EnableAreaBandwidth bool `json:"enable_area_bandwidth"`

	// 是否支持95转按需。
	EnableChange95 bool `json:"enable_change_95"`

	// 是否支持多SKU产品功能。
	EnableSpecCode *bool `json:"enable_spec_code,omitempty"`
}

func (o ListGlobalConnectionBandwidthConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthConfigs struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthConfigs", string(data)}, " ")
}

type ListGlobalConnectionBandwidthConfigsChargeMode struct {
	value string
}

type ListGlobalConnectionBandwidthConfigsChargeModeEnum struct {
	BWD     ListGlobalConnectionBandwidthConfigsChargeMode
	E_95    ListGlobalConnectionBandwidthConfigsChargeMode
	E_95AVR ListGlobalConnectionBandwidthConfigsChargeMode
}

func GetListGlobalConnectionBandwidthConfigsChargeModeEnum() ListGlobalConnectionBandwidthConfigsChargeModeEnum {
	return ListGlobalConnectionBandwidthConfigsChargeModeEnum{
		BWD: ListGlobalConnectionBandwidthConfigsChargeMode{
			value: "bwd",
		},
		E_95: ListGlobalConnectionBandwidthConfigsChargeMode{
			value: "95",
		},
		E_95AVR: ListGlobalConnectionBandwidthConfigsChargeMode{
			value: "95avr",
		},
	}
}

func (c ListGlobalConnectionBandwidthConfigsChargeMode) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthConfigsChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthConfigsChargeMode) UnmarshalJSON(b []byte) error {
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

type ListGlobalConnectionBandwidthConfigsServices struct {
	value string
}

type ListGlobalConnectionBandwidthConfigsServicesEnum struct {
	CC   ListGlobalConnectionBandwidthConfigsServices
	GEIP ListGlobalConnectionBandwidthConfigsServices
	GCN  ListGlobalConnectionBandwidthConfigsServices
	GSN  ListGlobalConnectionBandwidthConfigsServices
}

func GetListGlobalConnectionBandwidthConfigsServicesEnum() ListGlobalConnectionBandwidthConfigsServicesEnum {
	return ListGlobalConnectionBandwidthConfigsServicesEnum{
		CC: ListGlobalConnectionBandwidthConfigsServices{
			value: "CC",
		},
		GEIP: ListGlobalConnectionBandwidthConfigsServices{
			value: "GEIP",
		},
		GCN: ListGlobalConnectionBandwidthConfigsServices{
			value: "GCN",
		},
		GSN: ListGlobalConnectionBandwidthConfigsServices{
			value: "GSN",
		},
	}
}

func (c ListGlobalConnectionBandwidthConfigsServices) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthConfigsServices) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthConfigsServices) UnmarshalJSON(b []byte) error {
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

type ListGlobalConnectionBandwidthConfigsGcbType struct {
	value string
}

type ListGlobalConnectionBandwidthConfigsGcbTypeEnum struct {
	REGION   ListGlobalConnectionBandwidthConfigsGcbType
	SUB_AREA ListGlobalConnectionBandwidthConfigsGcbType
	AREA     ListGlobalConnectionBandwidthConfigsGcbType
	TRS_AREA ListGlobalConnectionBandwidthConfigsGcbType
}

func GetListGlobalConnectionBandwidthConfigsGcbTypeEnum() ListGlobalConnectionBandwidthConfigsGcbTypeEnum {
	return ListGlobalConnectionBandwidthConfigsGcbTypeEnum{
		REGION: ListGlobalConnectionBandwidthConfigsGcbType{
			value: "Region",
		},
		SUB_AREA: ListGlobalConnectionBandwidthConfigsGcbType{
			value: "SubArea",
		},
		AREA: ListGlobalConnectionBandwidthConfigsGcbType{
			value: "Area",
		},
		TRS_AREA: ListGlobalConnectionBandwidthConfigsGcbType{
			value: "TrsArea",
		},
	}
}

func (c ListGlobalConnectionBandwidthConfigsGcbType) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthConfigsGcbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthConfigsGcbType) UnmarshalJSON(b []byte) error {
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

type ListGlobalConnectionBandwidthConfigsSlaLevel struct {
	value string
}

type ListGlobalConnectionBandwidthConfigsSlaLevelEnum struct {
	PT ListGlobalConnectionBandwidthConfigsSlaLevel
	AU ListGlobalConnectionBandwidthConfigsSlaLevel
	AG ListGlobalConnectionBandwidthConfigsSlaLevel
}

func GetListGlobalConnectionBandwidthConfigsSlaLevelEnum() ListGlobalConnectionBandwidthConfigsSlaLevelEnum {
	return ListGlobalConnectionBandwidthConfigsSlaLevelEnum{
		PT: ListGlobalConnectionBandwidthConfigsSlaLevel{
			value: "Pt",
		},
		AU: ListGlobalConnectionBandwidthConfigsSlaLevel{
			value: "Au",
		},
		AG: ListGlobalConnectionBandwidthConfigsSlaLevel{
			value: "Ag",
		},
	}
}

func (c ListGlobalConnectionBandwidthConfigsSlaLevel) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthConfigsSlaLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthConfigsSlaLevel) UnmarshalJSON(b []byte) error {
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

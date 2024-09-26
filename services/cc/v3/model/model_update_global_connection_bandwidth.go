package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateGlobalConnectionBandwidth 更新全域互联带宽的详细信息
type UpdateGlobalConnectionBandwidth struct {

	// 实例名字。
	Name *string `json:"name,omitempty"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 功能说明：全域互联带宽实例中的带宽值大小，单位Mbit/s。 取值范围：2-300Mbit/s
	Size *int32 `json:"size,omitempty"`

	// 功能说明：描述计费类型，描述可选计费类型。默认开放按带宽计费，传统95计费租户白名单控制。 取值范围：     bwd: 按带宽计费     95: 按传统型95计费     95avr: 按传统型日95计费
	ChargeMode *UpdateGlobalConnectionBandwidthChargeMode `json:"charge_mode,omitempty"`

	// 功能说明：描述网络等级，从高到低分为铂金、金、银。默认金，其余租户白名单控制。 - Pt: 铂金 - Au: 金 - Ag: 银
	SlaLevel *UpdateGlobalConnectionBandwidthSlaLevel `json:"sla_level,omitempty"`

	// 功能说明：绑定的服务类型。实例类型： - CC: 云连接 - GEIP: 全域弹性公网IP - GCN: 中心网络 - GSN: 分支网络 - ALL: 所有实例类型
	BindingService *UpdateGlobalConnectionBandwidthBindingService `json:"binding_service,omitempty"`

	// 功能说明：线路规格编码UUID。
	SpecCodeId *string `json:"spec_code_id,omitempty"`
}

func (o UpdateGlobalConnectionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalConnectionBandwidth struct{}"
	}

	return strings.Join([]string{"UpdateGlobalConnectionBandwidth", string(data)}, " ")
}

type UpdateGlobalConnectionBandwidthChargeMode struct {
	value string
}

type UpdateGlobalConnectionBandwidthChargeModeEnum struct {
	BWD     UpdateGlobalConnectionBandwidthChargeMode
	E_95    UpdateGlobalConnectionBandwidthChargeMode
	E_95AVR UpdateGlobalConnectionBandwidthChargeMode
}

func GetUpdateGlobalConnectionBandwidthChargeModeEnum() UpdateGlobalConnectionBandwidthChargeModeEnum {
	return UpdateGlobalConnectionBandwidthChargeModeEnum{
		BWD: UpdateGlobalConnectionBandwidthChargeMode{
			value: "bwd",
		},
		E_95: UpdateGlobalConnectionBandwidthChargeMode{
			value: "95",
		},
		E_95AVR: UpdateGlobalConnectionBandwidthChargeMode{
			value: "95avr",
		},
	}
}

func (c UpdateGlobalConnectionBandwidthChargeMode) Value() string {
	return c.value
}

func (c UpdateGlobalConnectionBandwidthChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateGlobalConnectionBandwidthChargeMode) UnmarshalJSON(b []byte) error {
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

type UpdateGlobalConnectionBandwidthSlaLevel struct {
	value string
}

type UpdateGlobalConnectionBandwidthSlaLevelEnum struct {
	PT UpdateGlobalConnectionBandwidthSlaLevel
	AU UpdateGlobalConnectionBandwidthSlaLevel
	AG UpdateGlobalConnectionBandwidthSlaLevel
}

func GetUpdateGlobalConnectionBandwidthSlaLevelEnum() UpdateGlobalConnectionBandwidthSlaLevelEnum {
	return UpdateGlobalConnectionBandwidthSlaLevelEnum{
		PT: UpdateGlobalConnectionBandwidthSlaLevel{
			value: "Pt",
		},
		AU: UpdateGlobalConnectionBandwidthSlaLevel{
			value: "Au",
		},
		AG: UpdateGlobalConnectionBandwidthSlaLevel{
			value: "Ag",
		},
	}
}

func (c UpdateGlobalConnectionBandwidthSlaLevel) Value() string {
	return c.value
}

func (c UpdateGlobalConnectionBandwidthSlaLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateGlobalConnectionBandwidthSlaLevel) UnmarshalJSON(b []byte) error {
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

type UpdateGlobalConnectionBandwidthBindingService struct {
	value string
}

type UpdateGlobalConnectionBandwidthBindingServiceEnum struct {
	CC   UpdateGlobalConnectionBandwidthBindingService
	GEIP UpdateGlobalConnectionBandwidthBindingService
	GCN  UpdateGlobalConnectionBandwidthBindingService
	GSN  UpdateGlobalConnectionBandwidthBindingService
	ALL  UpdateGlobalConnectionBandwidthBindingService
}

func GetUpdateGlobalConnectionBandwidthBindingServiceEnum() UpdateGlobalConnectionBandwidthBindingServiceEnum {
	return UpdateGlobalConnectionBandwidthBindingServiceEnum{
		CC: UpdateGlobalConnectionBandwidthBindingService{
			value: "CC",
		},
		GEIP: UpdateGlobalConnectionBandwidthBindingService{
			value: "GEIP",
		},
		GCN: UpdateGlobalConnectionBandwidthBindingService{
			value: "GCN",
		},
		GSN: UpdateGlobalConnectionBandwidthBindingService{
			value: "GSN",
		},
		ALL: UpdateGlobalConnectionBandwidthBindingService{
			value: "ALL",
		},
	}
}

func (c UpdateGlobalConnectionBandwidthBindingService) Value() string {
	return c.value
}

func (c UpdateGlobalConnectionBandwidthBindingService) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateGlobalConnectionBandwidthBindingService) UnmarshalJSON(b []byte) error {
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

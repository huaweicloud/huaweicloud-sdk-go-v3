package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateGlobalConnectionBandwidth 创建全域互联带宽实例的详细信息。
type CreateGlobalConnectionBandwidth struct {

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 功能说明：全域互联带宽是否跨境，判断依据：带宽是否涉及从中国大陆到其他国家。 取值范围：True：跨境；False：非跨境
	Bordercross bool `json:"bordercross"`

	// 功能说明：描述带宽类型，对应地理区间的城域、区域、大区、跨区四级： - TrsArea: 跨区带宽 - Area: 大区带宽 - SubArea: 区域带宽 - Region: 城域带宽
	Type CreateGlobalConnectionBandwidthType `json:"type"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例标签。
	Tags *[]Tag `json:"tags,omitempty"`

	// 功能说明：描述计费类型，描述可选计费类型。默认开放按带宽计费，传统95计费租户白名单控制。 取值范围：     bwd: 按带宽计费     95: 按传统型95计费     95avr: 按传统型日95计费
	ChargeMode CreateGlobalConnectionBandwidthChargeMode `json:"charge_mode"`

	// 功能说明：全域互联带宽实例中的带宽值大小，单位Mbit/s。 取值范围：2-300Mbit/s
	Size int32 `json:"size"`

	// 功能说明：描述网络等级，从高到低分为铂金、金、银。默认金，其余租户白名单控制。 - Pt: 铂金 - Au: 金 - Ag: 银
	SlaLevel *CreateGlobalConnectionBandwidthSlaLevel `json:"sla_level,omitempty"`

	// 功能说明：本端接入点，配合remote_area信息描述带宽实例应用的范围。 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点），站点编码通过接口获取，带宽类型为Region可不传，其他类型必传
	LocalArea *string `json:"local_area,omitempty"`

	// 功能说明：远端接入点，配合local_area信息描述带宽实例应用的范围。 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点），站点编码通过接口获取，带宽类型为Region可不传，其他类型必传
	RemoteArea *string `json:"remote_area,omitempty"`

	// 功能说明：线路规格编码UUID。
	SpecCodeId *string `json:"spec_code_id,omitempty"`
}

func (o CreateGlobalConnectionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalConnectionBandwidth struct{}"
	}

	return strings.Join([]string{"CreateGlobalConnectionBandwidth", string(data)}, " ")
}

type CreateGlobalConnectionBandwidthType struct {
	value string
}

type CreateGlobalConnectionBandwidthTypeEnum struct {
	REGION   CreateGlobalConnectionBandwidthType
	SUB_AREA CreateGlobalConnectionBandwidthType
	AREA     CreateGlobalConnectionBandwidthType
	TRS_AREA CreateGlobalConnectionBandwidthType
}

func GetCreateGlobalConnectionBandwidthTypeEnum() CreateGlobalConnectionBandwidthTypeEnum {
	return CreateGlobalConnectionBandwidthTypeEnum{
		REGION: CreateGlobalConnectionBandwidthType{
			value: "Region",
		},
		SUB_AREA: CreateGlobalConnectionBandwidthType{
			value: "SubArea",
		},
		AREA: CreateGlobalConnectionBandwidthType{
			value: "Area",
		},
		TRS_AREA: CreateGlobalConnectionBandwidthType{
			value: "TrsArea",
		},
	}
}

func (c CreateGlobalConnectionBandwidthType) Value() string {
	return c.value
}

func (c CreateGlobalConnectionBandwidthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGlobalConnectionBandwidthType) UnmarshalJSON(b []byte) error {
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

type CreateGlobalConnectionBandwidthChargeMode struct {
	value string
}

type CreateGlobalConnectionBandwidthChargeModeEnum struct {
	BWD     CreateGlobalConnectionBandwidthChargeMode
	E_95    CreateGlobalConnectionBandwidthChargeMode
	E_95AVR CreateGlobalConnectionBandwidthChargeMode
}

func GetCreateGlobalConnectionBandwidthChargeModeEnum() CreateGlobalConnectionBandwidthChargeModeEnum {
	return CreateGlobalConnectionBandwidthChargeModeEnum{
		BWD: CreateGlobalConnectionBandwidthChargeMode{
			value: "bwd",
		},
		E_95: CreateGlobalConnectionBandwidthChargeMode{
			value: "95",
		},
		E_95AVR: CreateGlobalConnectionBandwidthChargeMode{
			value: "95avr",
		},
	}
}

func (c CreateGlobalConnectionBandwidthChargeMode) Value() string {
	return c.value
}

func (c CreateGlobalConnectionBandwidthChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGlobalConnectionBandwidthChargeMode) UnmarshalJSON(b []byte) error {
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

type CreateGlobalConnectionBandwidthSlaLevel struct {
	value string
}

type CreateGlobalConnectionBandwidthSlaLevelEnum struct {
	PT CreateGlobalConnectionBandwidthSlaLevel
	AU CreateGlobalConnectionBandwidthSlaLevel
	AG CreateGlobalConnectionBandwidthSlaLevel
}

func GetCreateGlobalConnectionBandwidthSlaLevelEnum() CreateGlobalConnectionBandwidthSlaLevelEnum {
	return CreateGlobalConnectionBandwidthSlaLevelEnum{
		PT: CreateGlobalConnectionBandwidthSlaLevel{
			value: "Pt",
		},
		AU: CreateGlobalConnectionBandwidthSlaLevel{
			value: "Au",
		},
		AG: CreateGlobalConnectionBandwidthSlaLevel{
			value: "Ag",
		},
	}
}

func (c CreateGlobalConnectionBandwidthSlaLevel) Value() string {
	return c.value
}

func (c CreateGlobalConnectionBandwidthSlaLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGlobalConnectionBandwidthSlaLevel) UnmarshalJSON(b []byte) error {
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

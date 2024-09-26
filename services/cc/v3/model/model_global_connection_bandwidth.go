package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// GlobalConnectionBandwidth 全域互联带宽实例的详细信息。
type GlobalConnectionBandwidth struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	// 功能说明：全域互联带宽是否跨境，判断依据：带宽是否涉及从中国大陆到其他国家。 取值范围：True：跨境；False：非跨境
	Bordercross bool `json:"bordercross"`

	// 功能说明：描述带宽类型，对应地理区间的城域、区域、大区、跨区四级： - TrsArea: 跨区带宽 - Area: 大区带宽 - SubArea: 区域带宽 - Region: 城域带宽
	Type GlobalConnectionBandwidthType `json:"type"`

	// 功能说明：绑定的服务类型。实例类型： - CC: 云连接 - GEIP: 全域弹性公网IP - GCN: 中心网络 - GSN: 分支网络 - ALL: 所有实例类型
	BindingService *GlobalConnectionBandwidthBindingService `json:"binding_service,omitempty"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 功能说明：描述计费类型，描述可选计费类型。默认开放按带宽计费，传统95计费租户白名单控制。 取值范围：     bwd: 按带宽计费     95: 按传统型95计费     95avr: 按传统型日95计费
	ChargeMode GlobalConnectionBandwidthChargeMode `json:"charge_mode"`

	// 功能说明：全域互联带宽实例中的带宽值大小，单位Mbit/s。 取值范围：2-300Mbit/s
	Size int32 `json:"size"`

	// 功能说明：描述网络等级，从高到低分为铂金、金、银。默认金，其余租户白名单控制。 - Pt: 铂金 - Au: 金 - Ag: 银
	SlaLevel *GlobalConnectionBandwidthSlaLevel `json:"sla_level,omitempty"`

	// 功能说明：本端接入点的中英文名。通过HEADER里面的x-language控制，默认英文，zh-cn返回中文。
	LocalArea *string `json:"local_area,omitempty"`

	// 功能说明：远端接入点的中英文名。通过HEADER里面的x-language控制，默认英文，zh-cn返回中文。
	RemoteArea *string `json:"remote_area,omitempty"`

	// 功能说明：本端接入点的编码。
	LocalSiteCode *string `json:"local_site_code,omitempty"`

	// 功能说明：远端接入点的编码。
	RemoteSiteCode *string `json:"remote_site_code,omitempty"`

	// 功能说明: 全域互联带宽状态。 取值范围：     NORMAL-正常     FREEZED-冻结状态
	AdminState *GlobalConnectionBandwidthAdminState `json:"admin_state,omitempty"`

	// 功能说明: 全域互联带宽是否冻结。 取值范围：     true-冻结     false-非冻结
	Frozen *bool `json:"frozen,omitempty"`

	// 功能说明：线路规格编码UUID。
	SpecCodeId *string `json:"spec_code_id,omitempty"`

	// 实例标签。
	Tags *[]Tag `json:"tags,omitempty"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 功能说明: 全域互联带宽是否支持绑定多实例。 取值范围：     true-支持     false-不支持
	EnableShare *bool `json:"enable_share,omitempty"`

	// 功能说明: 全域互联带宽绑定实例列表。
	Instances *[]GlobalConnectionBandwidthAssociatedInstance `json:"instances,omitempty"`
}

func (o GlobalConnectionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalConnectionBandwidth struct{}"
	}

	return strings.Join([]string{"GlobalConnectionBandwidth", string(data)}, " ")
}

type GlobalConnectionBandwidthType struct {
	value string
}

type GlobalConnectionBandwidthTypeEnum struct {
	REGION   GlobalConnectionBandwidthType
	SUB_AREA GlobalConnectionBandwidthType
	AREA     GlobalConnectionBandwidthType
	TRS_AREA GlobalConnectionBandwidthType
}

func GetGlobalConnectionBandwidthTypeEnum() GlobalConnectionBandwidthTypeEnum {
	return GlobalConnectionBandwidthTypeEnum{
		REGION: GlobalConnectionBandwidthType{
			value: "Region",
		},
		SUB_AREA: GlobalConnectionBandwidthType{
			value: "SubArea",
		},
		AREA: GlobalConnectionBandwidthType{
			value: "Area",
		},
		TRS_AREA: GlobalConnectionBandwidthType{
			value: "TrsArea",
		},
	}
}

func (c GlobalConnectionBandwidthType) Value() string {
	return c.value
}

func (c GlobalConnectionBandwidthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalConnectionBandwidthType) UnmarshalJSON(b []byte) error {
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

type GlobalConnectionBandwidthBindingService struct {
	value string
}

type GlobalConnectionBandwidthBindingServiceEnum struct {
	CC   GlobalConnectionBandwidthBindingService
	GEIP GlobalConnectionBandwidthBindingService
	GCN  GlobalConnectionBandwidthBindingService
	GSN  GlobalConnectionBandwidthBindingService
	ALL  GlobalConnectionBandwidthBindingService
}

func GetGlobalConnectionBandwidthBindingServiceEnum() GlobalConnectionBandwidthBindingServiceEnum {
	return GlobalConnectionBandwidthBindingServiceEnum{
		CC: GlobalConnectionBandwidthBindingService{
			value: "CC",
		},
		GEIP: GlobalConnectionBandwidthBindingService{
			value: "GEIP",
		},
		GCN: GlobalConnectionBandwidthBindingService{
			value: "GCN",
		},
		GSN: GlobalConnectionBandwidthBindingService{
			value: "GSN",
		},
		ALL: GlobalConnectionBandwidthBindingService{
			value: "ALL",
		},
	}
}

func (c GlobalConnectionBandwidthBindingService) Value() string {
	return c.value
}

func (c GlobalConnectionBandwidthBindingService) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalConnectionBandwidthBindingService) UnmarshalJSON(b []byte) error {
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

type GlobalConnectionBandwidthChargeMode struct {
	value string
}

type GlobalConnectionBandwidthChargeModeEnum struct {
	BWD     GlobalConnectionBandwidthChargeMode
	E_95    GlobalConnectionBandwidthChargeMode
	E_95AVR GlobalConnectionBandwidthChargeMode
}

func GetGlobalConnectionBandwidthChargeModeEnum() GlobalConnectionBandwidthChargeModeEnum {
	return GlobalConnectionBandwidthChargeModeEnum{
		BWD: GlobalConnectionBandwidthChargeMode{
			value: "bwd",
		},
		E_95: GlobalConnectionBandwidthChargeMode{
			value: "95",
		},
		E_95AVR: GlobalConnectionBandwidthChargeMode{
			value: "95avr",
		},
	}
}

func (c GlobalConnectionBandwidthChargeMode) Value() string {
	return c.value
}

func (c GlobalConnectionBandwidthChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalConnectionBandwidthChargeMode) UnmarshalJSON(b []byte) error {
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

type GlobalConnectionBandwidthSlaLevel struct {
	value string
}

type GlobalConnectionBandwidthSlaLevelEnum struct {
	PT GlobalConnectionBandwidthSlaLevel
	AU GlobalConnectionBandwidthSlaLevel
	AG GlobalConnectionBandwidthSlaLevel
}

func GetGlobalConnectionBandwidthSlaLevelEnum() GlobalConnectionBandwidthSlaLevelEnum {
	return GlobalConnectionBandwidthSlaLevelEnum{
		PT: GlobalConnectionBandwidthSlaLevel{
			value: "Pt",
		},
		AU: GlobalConnectionBandwidthSlaLevel{
			value: "Au",
		},
		AG: GlobalConnectionBandwidthSlaLevel{
			value: "Ag",
		},
	}
}

func (c GlobalConnectionBandwidthSlaLevel) Value() string {
	return c.value
}

func (c GlobalConnectionBandwidthSlaLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalConnectionBandwidthSlaLevel) UnmarshalJSON(b []byte) error {
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

type GlobalConnectionBandwidthAdminState struct {
	value string
}

type GlobalConnectionBandwidthAdminStateEnum struct {
	NORMAL  GlobalConnectionBandwidthAdminState
	FREEZED GlobalConnectionBandwidthAdminState
}

func GetGlobalConnectionBandwidthAdminStateEnum() GlobalConnectionBandwidthAdminStateEnum {
	return GlobalConnectionBandwidthAdminStateEnum{
		NORMAL: GlobalConnectionBandwidthAdminState{
			value: "NORMAL",
		},
		FREEZED: GlobalConnectionBandwidthAdminState{
			value: "FREEZED",
		},
	}
}

func (c GlobalConnectionBandwidthAdminState) Value() string {
	return c.value
}

func (c GlobalConnectionBandwidthAdminState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalConnectionBandwidthAdminState) UnmarshalJSON(b []byte) error {
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

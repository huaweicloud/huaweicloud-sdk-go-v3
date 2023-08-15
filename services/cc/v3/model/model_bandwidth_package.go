package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// BandwidthPackage 带宽包实例。
type BandwidthPackage struct {

	// 带宽包实例的ID。
	Id *string `json:"id,omitempty"`

	// 带宽包实例的名字。
	Name *string `json:"name,omitempty"`

	// 带宽包实例的描述。
	Description *string `json:"description,omitempty"`

	// 帐号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 带宽包实例的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 带宽包实例的状态。ACTIVE表示状态
	Status *BandwidthPackageStatus `json:"status,omitempty"`

	// 带宽包实例的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 带宽包实例的更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 带宽包实例的管理状态。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 带宽包实例的订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 带宽包实例的产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 带宽包实例的规格编码。 bandwidth.aftoela：大陆站+国际站南非-拉美东 bandwidth.aftonla：大陆站+国际站南非-拉美北 bandwidth.aftowla：大陆站+国际站南非-拉美西 bandwidth.aptoaf：国际站亚太-南非 bandwidth.aptoap：国际站亚太-亚太 bandwidth.aptoela：大陆站+国际站亚太-拉美东 bandwidth.aptonla：大陆站+国际站亚太-拉美北 bandwidth.aptowla：大陆站+国际站亚太-拉美西 bandwidth.cmtoaf：国际站中国大陆-南非 bandwidth.cmtoap：国际站中国大陆-亚太 bandwidth.cmtocm：国际站中国大陆-中国大陆 bandwidth.cmtoela：大陆站+国际站中国大陆-拉美东 bandwidth.cmtonla：大陆站+国际站中国大陆-拉美北 bandwidth.cmtowla：大陆站+国际站中国大陆-拉美西 bandwidth.elatoela：大陆站+国际站拉美东-拉美东 bandwidth.elatonla：大陆站+国际站拉美东-拉美北 bandwidth.wlatoela：大陆站+国际站拉美西-拉美东 bandwidth.wlatonla：大陆站+国际站拉美西-拉美北 bandwidth.wlatowla：大陆站+国际站拉美西-拉美西
	SpecCode *BandwidthPackageSpecCode `json:"spec_code,omitempty"`

	// 带宽包实例在大陆站或国际站的计费方式。 1：大陆站包周期 2：国际站包周期 3：大陆站按需计费 4：国际站按需计费 5：大陆站按95方式计费 6：国际站按95方式计费
	BillingMode *BandwidthPackageBillingMode `json:"billing_mode,omitempty"`

	// 带宽包实例的计费方式。 bandwidth是按带宽计费。
	ChargeMode *BandwidthPackageChargeMode `json:"charge_mode,omitempty"`

	// 带宽包实例中的带宽值。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 带宽包实例绑定的资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 带宽包实例绑定的资源类型。 cloud_connection: 云连接实例。
	ResourceType *BandwidthPackageResourceType `json:"resource_type,omitempty"`

	// 本端大区。 云连接支持的大区有： | Chinese-Mainland | 中国大陆 | | Asia-Pacific | 亚太 | | Africa | 非洲 | | Western-Latin-America | 拉美西 | | Eastern-Latin-America | 拉美东 | | Northern-Latin-America | 拉美北 |
	LocalAreaId *BandwidthPackageLocalAreaId `json:"local_area_id,omitempty"`

	// 远端大区。 云连接支持的大区有： | Chinese-Mainland | 中国大陆 | | Asia-Pacific | 亚太 | | Africa | 非洲 | | Western-Latin-America | 拉美西 | | Eastern-Latin-America | 拉美东 | | Northern-Latin-America | 拉美北 |
	RemoteAreaId *BandwidthPackageRemoteAreaId `json:"remote_area_id,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 互通类型: - Area: 大区互通 - Region: 城域互通
	InterflowMode *BandwidthPackageInterflowMode `json:"interflow_mode,omitempty"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o BandwidthPackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthPackage struct{}"
	}

	return strings.Join([]string{"BandwidthPackage", string(data)}, " ")
}

type BandwidthPackageStatus struct {
	value string
}

type BandwidthPackageStatusEnum struct {
	ACTIVE BandwidthPackageStatus
}

func GetBandwidthPackageStatusEnum() BandwidthPackageStatusEnum {
	return BandwidthPackageStatusEnum{
		ACTIVE: BandwidthPackageStatus{
			value: "ACTIVE",
		},
	}
}

func (c BandwidthPackageStatus) Value() string {
	return c.value
}

func (c BandwidthPackageStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthPackageStatus) UnmarshalJSON(b []byte) error {
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

type BandwidthPackageSpecCode struct {
	value string
}

type BandwidthPackageSpecCodeEnum struct {
	BANDWIDTH_AFTOELA  BandwidthPackageSpecCode
	BANDWIDTH_AFTONLA  BandwidthPackageSpecCode
	BANDWIDTH_AFTOWLA  BandwidthPackageSpecCode
	BANDWIDTH_APTOAF   BandwidthPackageSpecCode
	BANDWIDTH_APTOAP   BandwidthPackageSpecCode
	BANDWIDTH_APTOELA  BandwidthPackageSpecCode
	BANDWIDTH_APTONLA  BandwidthPackageSpecCode
	BANDWIDTH_APTOWLA  BandwidthPackageSpecCode
	BANDWIDTH_CMTOAF   BandwidthPackageSpecCode
	BANDWIDTH_CMTOAP   BandwidthPackageSpecCode
	BANDWIDTH_CMTOCM   BandwidthPackageSpecCode
	BANDWIDTH_CMTOELA  BandwidthPackageSpecCode
	BANDWIDTH_CMTONLA  BandwidthPackageSpecCode
	BANDWIDTH_CMTOWLA  BandwidthPackageSpecCode
	BANDWIDTH_ELATOELA BandwidthPackageSpecCode
	BANDWIDTH_ELATONLA BandwidthPackageSpecCode
	BANDWIDTH_WLATOELA BandwidthPackageSpecCode
	BANDWIDTH_WLATONLA BandwidthPackageSpecCode
	BANDWIDTH_WLATOWLA BandwidthPackageSpecCode
}

func GetBandwidthPackageSpecCodeEnum() BandwidthPackageSpecCodeEnum {
	return BandwidthPackageSpecCodeEnum{
		BANDWIDTH_AFTOELA: BandwidthPackageSpecCode{
			value: "bandwidth.aftoela",
		},
		BANDWIDTH_AFTONLA: BandwidthPackageSpecCode{
			value: "bandwidth.aftonla",
		},
		BANDWIDTH_AFTOWLA: BandwidthPackageSpecCode{
			value: "bandwidth.aftowla",
		},
		BANDWIDTH_APTOAF: BandwidthPackageSpecCode{
			value: "bandwidth.aptoaf",
		},
		BANDWIDTH_APTOAP: BandwidthPackageSpecCode{
			value: "bandwidth.aptoap",
		},
		BANDWIDTH_APTOELA: BandwidthPackageSpecCode{
			value: "bandwidth.aptoela",
		},
		BANDWIDTH_APTONLA: BandwidthPackageSpecCode{
			value: "bandwidth.aptonla",
		},
		BANDWIDTH_APTOWLA: BandwidthPackageSpecCode{
			value: "bandwidth.aptowla",
		},
		BANDWIDTH_CMTOAF: BandwidthPackageSpecCode{
			value: "bandwidth.cmtoaf",
		},
		BANDWIDTH_CMTOAP: BandwidthPackageSpecCode{
			value: "bandwidth.cmtoap",
		},
		BANDWIDTH_CMTOCM: BandwidthPackageSpecCode{
			value: "bandwidth.cmtocm",
		},
		BANDWIDTH_CMTOELA: BandwidthPackageSpecCode{
			value: "bandwidth.cmtoela",
		},
		BANDWIDTH_CMTONLA: BandwidthPackageSpecCode{
			value: "bandwidth.cmtonla",
		},
		BANDWIDTH_CMTOWLA: BandwidthPackageSpecCode{
			value: "bandwidth.cmtowla",
		},
		BANDWIDTH_ELATOELA: BandwidthPackageSpecCode{
			value: "bandwidth.elatoela",
		},
		BANDWIDTH_ELATONLA: BandwidthPackageSpecCode{
			value: "bandwidth.elatonla",
		},
		BANDWIDTH_WLATOELA: BandwidthPackageSpecCode{
			value: "bandwidth.wlatoela",
		},
		BANDWIDTH_WLATONLA: BandwidthPackageSpecCode{
			value: "bandwidth.wlatonla",
		},
		BANDWIDTH_WLATOWLA: BandwidthPackageSpecCode{
			value: "bandwidth.wlatowla",
		},
	}
}

func (c BandwidthPackageSpecCode) Value() string {
	return c.value
}

func (c BandwidthPackageSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthPackageSpecCode) UnmarshalJSON(b []byte) error {
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

type BandwidthPackageBillingMode struct {
	value string
}

type BandwidthPackageBillingModeEnum struct {
	E_1 BandwidthPackageBillingMode
	E_2 BandwidthPackageBillingMode
	E_3 BandwidthPackageBillingMode
	E_4 BandwidthPackageBillingMode
	E_5 BandwidthPackageBillingMode
	E_6 BandwidthPackageBillingMode
}

func GetBandwidthPackageBillingModeEnum() BandwidthPackageBillingModeEnum {
	return BandwidthPackageBillingModeEnum{
		E_1: BandwidthPackageBillingMode{
			value: "1",
		},
		E_2: BandwidthPackageBillingMode{
			value: "2",
		},
		E_3: BandwidthPackageBillingMode{
			value: "3",
		},
		E_4: BandwidthPackageBillingMode{
			value: "4",
		},
		E_5: BandwidthPackageBillingMode{
			value: "5",
		},
		E_6: BandwidthPackageBillingMode{
			value: "6",
		},
	}
}

func (c BandwidthPackageBillingMode) Value() string {
	return c.value
}

func (c BandwidthPackageBillingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthPackageBillingMode) UnmarshalJSON(b []byte) error {
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

type BandwidthPackageChargeMode struct {
	value string
}

type BandwidthPackageChargeModeEnum struct {
	BANDWIDTH BandwidthPackageChargeMode
}

func GetBandwidthPackageChargeModeEnum() BandwidthPackageChargeModeEnum {
	return BandwidthPackageChargeModeEnum{
		BANDWIDTH: BandwidthPackageChargeMode{
			value: "bandwidth",
		},
	}
}

func (c BandwidthPackageChargeMode) Value() string {
	return c.value
}

func (c BandwidthPackageChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthPackageChargeMode) UnmarshalJSON(b []byte) error {
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

type BandwidthPackageResourceType struct {
	value string
}

type BandwidthPackageResourceTypeEnum struct {
	CLOUD_CONNECTION BandwidthPackageResourceType
}

func GetBandwidthPackageResourceTypeEnum() BandwidthPackageResourceTypeEnum {
	return BandwidthPackageResourceTypeEnum{
		CLOUD_CONNECTION: BandwidthPackageResourceType{
			value: "cloud_connection",
		},
	}
}

func (c BandwidthPackageResourceType) Value() string {
	return c.value
}

func (c BandwidthPackageResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthPackageResourceType) UnmarshalJSON(b []byte) error {
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

type BandwidthPackageLocalAreaId struct {
	value string
}

type BandwidthPackageLocalAreaIdEnum struct {
	CHINESE_MAINLAND       BandwidthPackageLocalAreaId
	ASIA_PACIFIC           BandwidthPackageLocalAreaId
	AFRICA                 BandwidthPackageLocalAreaId
	WESTERN_LATIN_AMERICA  BandwidthPackageLocalAreaId
	EASTERN_LATIN_AMERICA  BandwidthPackageLocalAreaId
	NORTHERN_LATIN_AMERICA BandwidthPackageLocalAreaId
}

func GetBandwidthPackageLocalAreaIdEnum() BandwidthPackageLocalAreaIdEnum {
	return BandwidthPackageLocalAreaIdEnum{
		CHINESE_MAINLAND: BandwidthPackageLocalAreaId{
			value: "Chinese-Mainland",
		},
		ASIA_PACIFIC: BandwidthPackageLocalAreaId{
			value: "Asia-Pacific",
		},
		AFRICA: BandwidthPackageLocalAreaId{
			value: "Africa",
		},
		WESTERN_LATIN_AMERICA: BandwidthPackageLocalAreaId{
			value: "Western-Latin-America",
		},
		EASTERN_LATIN_AMERICA: BandwidthPackageLocalAreaId{
			value: "Eastern-Latin-America",
		},
		NORTHERN_LATIN_AMERICA: BandwidthPackageLocalAreaId{
			value: "Northern-Latin-America",
		},
	}
}

func (c BandwidthPackageLocalAreaId) Value() string {
	return c.value
}

func (c BandwidthPackageLocalAreaId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthPackageLocalAreaId) UnmarshalJSON(b []byte) error {
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

type BandwidthPackageRemoteAreaId struct {
	value string
}

type BandwidthPackageRemoteAreaIdEnum struct {
	CHINESE_MAINLAND       BandwidthPackageRemoteAreaId
	ASIA_PACIFIC           BandwidthPackageRemoteAreaId
	AFRICA                 BandwidthPackageRemoteAreaId
	WESTERN_LATIN_AMERICA  BandwidthPackageRemoteAreaId
	EASTERN_LATIN_AMERICA  BandwidthPackageRemoteAreaId
	NORTHERN_LATIN_AMERICA BandwidthPackageRemoteAreaId
}

func GetBandwidthPackageRemoteAreaIdEnum() BandwidthPackageRemoteAreaIdEnum {
	return BandwidthPackageRemoteAreaIdEnum{
		CHINESE_MAINLAND: BandwidthPackageRemoteAreaId{
			value: "Chinese-Mainland",
		},
		ASIA_PACIFIC: BandwidthPackageRemoteAreaId{
			value: "Asia-Pacific",
		},
		AFRICA: BandwidthPackageRemoteAreaId{
			value: "Africa",
		},
		WESTERN_LATIN_AMERICA: BandwidthPackageRemoteAreaId{
			value: "Western-Latin-America",
		},
		EASTERN_LATIN_AMERICA: BandwidthPackageRemoteAreaId{
			value: "Eastern-Latin-America",
		},
		NORTHERN_LATIN_AMERICA: BandwidthPackageRemoteAreaId{
			value: "Northern-Latin-America",
		},
	}
}

func (c BandwidthPackageRemoteAreaId) Value() string {
	return c.value
}

func (c BandwidthPackageRemoteAreaId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthPackageRemoteAreaId) UnmarshalJSON(b []byte) error {
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

type BandwidthPackageInterflowMode struct {
	value string
}

type BandwidthPackageInterflowModeEnum struct {
	AREA   BandwidthPackageInterflowMode
	REGION BandwidthPackageInterflowMode
}

func GetBandwidthPackageInterflowModeEnum() BandwidthPackageInterflowModeEnum {
	return BandwidthPackageInterflowModeEnum{
		AREA: BandwidthPackageInterflowMode{
			value: "Area",
		},
		REGION: BandwidthPackageInterflowMode{
			value: "Region",
		},
	}
}

func (c BandwidthPackageInterflowMode) Value() string {
	return c.value
}

func (c BandwidthPackageInterflowMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthPackageInterflowMode) UnmarshalJSON(b []byte) error {
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

package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// BandwidthPackage 带宽包实例。
type BandwidthPackage struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 带宽包实例绑定的资源ID。
	ResourceId string `json:"resource_id"`

	// 带宽包实例绑定的资源类型。 cloud_connection: 云连接实例。
	ResourceType BandwidthPackageResourceType `json:"resource_type"`

	LocalAreaId *AreaIdDef `json:"local_area_id"`

	RemoteAreaId *AreaIdDef `json:"remote_area_id"`

	// 带宽包实例的规格编码。 bandwidth.aftoela：大陆站+国际站南非-拉美东 bandwidth.aftonla：大陆站+国际站南非-拉美北 bandwidth.aftowla：大陆站+国际站南非-拉美西 bandwidth.aptoaf：国际站亚太-南非 bandwidth.aptoap：国际站亚太-亚太 bandwidth.aptoela：大陆站+国际站亚太-拉美东 bandwidth.aptonla：大陆站+国际站亚太-拉美北 bandwidth.aptowla：大陆站+国际站亚太-拉美西 bandwidth.cmtoaf：国际站中国大陆-南非 bandwidth.cmtoap：国际站中国大陆-亚太 bandwidth.cmtocm：国际站中国大陆-中国大陆 bandwidth.cmtoela：大陆站+国际站中国大陆-拉美东 bandwidth.cmtonla：大陆站+国际站中国大陆-拉美北 bandwidth.cmtowla：大陆站+国际站中国大陆-拉美西 bandwidth.elatoela：大陆站+国际站拉美东-拉美东 bandwidth.elatonla：大陆站+国际站拉美东-拉美北 bandwidth.wlatoela：大陆站+国际站拉美西-拉美东 bandwidth.wlatonla：大陆站+国际站拉美西-拉美北 bandwidth.wlatowla：大陆站+国际站拉美西-拉美西
	SpecCode string `json:"spec_code"`

	BillingMode *BillingModeEnum `json:"billing_mode"`

	// 实例标签。
	Tags *[]Tag `json:"tags,omitempty"`

	// 带宽包实例的状态。ACTIVE表示状态
	Status *BandwidthPackageStatus `json:"status,omitempty"`

	// 带宽包实例的管理状态。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 带宽包实例的订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 带宽包实例的产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 带宽包实例的计费方式。 bandwidth是按带宽计费。
	ChargeMode *BandwidthPackageChargeMode `json:"charge_mode,omitempty"`

	// 带宽包实例中的带宽值。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 互通类型: - Area: 大区互通 - Region: 城域互通
	InterflowMode *BandwidthPackageInterflowMode `json:"interflow_mode,omitempty"`
}

func (o BandwidthPackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthPackage struct{}"
	}

	return strings.Join([]string{"BandwidthPackage", string(data)}, " ")
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

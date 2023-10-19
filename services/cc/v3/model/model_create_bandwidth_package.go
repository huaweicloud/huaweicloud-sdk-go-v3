package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateBandwidthPackage 创建带宽包请求体。
type CreateBandwidthPackage struct {

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	LocalAreaId *AreaIdDef `json:"local_area_id"`

	RemoteAreaId *AreaIdDef `json:"remote_area_id"`

	// 带宽包实例的计费方式。 bandwidth是按带宽计费。
	ChargeMode CreateBandwidthPackageChargeMode `json:"charge_mode"`

	// 带宽包实例在大陆站或国际站的计费方式： - 3：大陆站按需计费 - 4：国际站按需计费 - 5：大陆站按95方式计费 - 6：国际站按95方式计费
	BillingMode CreateBandwidthPackageBillingMode `json:"billing_mode"`

	// 带宽包实例中的带宽值。
	Bandwidth int32 `json:"bandwidth"`

	// 带宽包实例绑定的资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 带宽包实例绑定的资源类型。  cloud_connection: 云连接实例。
	ResourceType *CreateBandwidthPackageResourceType `json:"resource_type,omitempty"`

	// 带宽包实例的规格编码。 bandwidth.aftoela：大陆站+国际站南非-拉美东 bandwidth.aftonla：大陆站+国际站南非-拉美北 bandwidth.aftowla：大陆站+国际站南非-拉美西 bandwidth.aptoaf：国际站亚太-南非 bandwidth.aptoap：国际站亚太-亚太 bandwidth.aptoela：大陆站+国际站亚太-拉美东 bandwidth.aptonla：大陆站+国际站亚太-拉美北 bandwidth.aptowla：大陆站+国际站亚太-拉美西 bandwidth.cmtoaf：国际站中国大陆-南非 bandwidth.cmtoap：国际站中国大陆-亚太 bandwidth.cmtocm：国际站中国大陆-中国大陆 bandwidth.cmtoela：大陆站+国际站中国大陆-拉美东 bandwidth.cmtonla：大陆站+国际站中国大陆-拉美北 bandwidth.cmtowla：大陆站+国际站中国大陆-拉美西 bandwidth.elatoela：大陆站+国际站拉美东-拉美东 bandwidth.elatonla：大陆站+国际站拉美东-拉美北 bandwidth.wlatoela：大陆站+国际站拉美西-拉美东 bandwidth.wlatonla：大陆站+国际站拉美西-拉美北 bandwidth.wlatowla：大陆站+国际站拉美西-拉美西
	SpecCode *string `json:"spec_code,omitempty"`

	// 互通类型: - Area: 大区互通 - Region: 城域互通
	InterflowMode *CreateBandwidthPackageInterflowMode `json:"interflow_mode,omitempty"`
}

func (o CreateBandwidthPackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBandwidthPackage struct{}"
	}

	return strings.Join([]string{"CreateBandwidthPackage", string(data)}, " ")
}

type CreateBandwidthPackageChargeMode struct {
	value string
}

type CreateBandwidthPackageChargeModeEnum struct {
	BANDWIDTH CreateBandwidthPackageChargeMode
}

func GetCreateBandwidthPackageChargeModeEnum() CreateBandwidthPackageChargeModeEnum {
	return CreateBandwidthPackageChargeModeEnum{
		BANDWIDTH: CreateBandwidthPackageChargeMode{
			value: "bandwidth",
		},
	}
}

func (c CreateBandwidthPackageChargeMode) Value() string {
	return c.value
}

func (c CreateBandwidthPackageChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBandwidthPackageChargeMode) UnmarshalJSON(b []byte) error {
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

type CreateBandwidthPackageBillingMode struct {
	value int32
}

type CreateBandwidthPackageBillingModeEnum struct {
	E_3 CreateBandwidthPackageBillingMode
	E_4 CreateBandwidthPackageBillingMode
	E_5 CreateBandwidthPackageBillingMode
	E_6 CreateBandwidthPackageBillingMode
}

func GetCreateBandwidthPackageBillingModeEnum() CreateBandwidthPackageBillingModeEnum {
	return CreateBandwidthPackageBillingModeEnum{
		E_3: CreateBandwidthPackageBillingMode{
			value: 3,
		}, E_4: CreateBandwidthPackageBillingMode{
			value: 4,
		}, E_5: CreateBandwidthPackageBillingMode{
			value: 5,
		}, E_6: CreateBandwidthPackageBillingMode{
			value: 6,
		},
	}
}

func (c CreateBandwidthPackageBillingMode) Value() int32 {
	return c.value
}

func (c CreateBandwidthPackageBillingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBandwidthPackageBillingMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateBandwidthPackageResourceType struct {
	value string
}

type CreateBandwidthPackageResourceTypeEnum struct {
	CLOUD_CONNECTION CreateBandwidthPackageResourceType
}

func GetCreateBandwidthPackageResourceTypeEnum() CreateBandwidthPackageResourceTypeEnum {
	return CreateBandwidthPackageResourceTypeEnum{
		CLOUD_CONNECTION: CreateBandwidthPackageResourceType{
			value: "cloud_connection",
		},
	}
}

func (c CreateBandwidthPackageResourceType) Value() string {
	return c.value
}

func (c CreateBandwidthPackageResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBandwidthPackageResourceType) UnmarshalJSON(b []byte) error {
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

type CreateBandwidthPackageInterflowMode struct {
	value string
}

type CreateBandwidthPackageInterflowModeEnum struct {
	AREA   CreateBandwidthPackageInterflowMode
	REGION CreateBandwidthPackageInterflowMode
}

func GetCreateBandwidthPackageInterflowModeEnum() CreateBandwidthPackageInterflowModeEnum {
	return CreateBandwidthPackageInterflowModeEnum{
		AREA: CreateBandwidthPackageInterflowMode{
			value: "Area",
		},
		REGION: CreateBandwidthPackageInterflowMode{
			value: "Region",
		},
	}
}

func (c CreateBandwidthPackageInterflowMode) Value() string {
	return c.value
}

func (c CreateBandwidthPackageInterflowMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBandwidthPackageInterflowMode) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SetupLandingZoneReqBody 设置Landing Zone的参数。
type SetupLandingZoneReqBody struct {

	// 管理员纳管账号创建Identity Center用户所用邮箱。
	IdentityStoreEmail *string `json:"identity_store_email,omitempty"`

	// 主区域。
	HomeRegion string `json:"home_region"`

	// 设置Landing Zone的类型。包括CREATE、REPAIR以及UPDATE。
	SetupLandingZoneActionType SetupLandingZoneReqBodySetupLandingZoneActionType `json:"setup_landing_zone_action_type"`

	// 当前纳管账号纳管的区域。
	RegionConfigurationList []RegionConfigurationList `json:"region_configuration_list"`

	// 是否设置Landing Zone的identity center。
	IdentityCenterStatus *SetupLandingZoneReqBodyIdentityCenterStatus `json:"identity_center_status,omitempty"`

	// 设置organization的类型。STANDARD和NON_STANDARD。
	OrganizationStructureType *SetupLandingZoneReqBodyOrganizationStructureType `json:"organization_structure_type,omitempty"`

	// 基础环境的纳管账号体系。
	OrganizationStructure []OrganizationStructureBaseLine `json:"organization_structure"`

	// 是否允许区域拒绝，默认false。
	DenyUngovernedRegions *bool `json:"deny_ungoverned_regions,omitempty"`

	// 是否允许设置组织汇聚。
	CloudTrailType *bool `json:"cloud_trail_type,omitempty"`

	// 加密字段。
	KmsKeyId *string `json:"kms_key_id,omitempty"`

	LoggingConfiguration *LoggingConfiguration `json:"logging_configuration"`

	// 基线版本
	BaselineVersion *string `json:"baseline_version,omitempty"`
}

func (o SetupLandingZoneReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetupLandingZoneReqBody struct{}"
	}

	return strings.Join([]string{"SetupLandingZoneReqBody", string(data)}, " ")
}

type SetupLandingZoneReqBodySetupLandingZoneActionType struct {
	value string
}

type SetupLandingZoneReqBodySetupLandingZoneActionTypeEnum struct {
	CREATE SetupLandingZoneReqBodySetupLandingZoneActionType
	UPDATE SetupLandingZoneReqBodySetupLandingZoneActionType
	REPAIR SetupLandingZoneReqBodySetupLandingZoneActionType
}

func GetSetupLandingZoneReqBodySetupLandingZoneActionTypeEnum() SetupLandingZoneReqBodySetupLandingZoneActionTypeEnum {
	return SetupLandingZoneReqBodySetupLandingZoneActionTypeEnum{
		CREATE: SetupLandingZoneReqBodySetupLandingZoneActionType{
			value: "CREATE",
		},
		UPDATE: SetupLandingZoneReqBodySetupLandingZoneActionType{
			value: "UPDATE",
		},
		REPAIR: SetupLandingZoneReqBodySetupLandingZoneActionType{
			value: "REPAIR",
		},
	}
}

func (c SetupLandingZoneReqBodySetupLandingZoneActionType) Value() string {
	return c.value
}

func (c SetupLandingZoneReqBodySetupLandingZoneActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetupLandingZoneReqBodySetupLandingZoneActionType) UnmarshalJSON(b []byte) error {
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

type SetupLandingZoneReqBodyIdentityCenterStatus struct {
	value string
}

type SetupLandingZoneReqBodyIdentityCenterStatusEnum struct {
	ENABLE  SetupLandingZoneReqBodyIdentityCenterStatus
	DISABLE SetupLandingZoneReqBodyIdentityCenterStatus
}

func GetSetupLandingZoneReqBodyIdentityCenterStatusEnum() SetupLandingZoneReqBodyIdentityCenterStatusEnum {
	return SetupLandingZoneReqBodyIdentityCenterStatusEnum{
		ENABLE: SetupLandingZoneReqBodyIdentityCenterStatus{
			value: "ENABLE",
		},
		DISABLE: SetupLandingZoneReqBodyIdentityCenterStatus{
			value: "DISABLE",
		},
	}
}

func (c SetupLandingZoneReqBodyIdentityCenterStatus) Value() string {
	return c.value
}

func (c SetupLandingZoneReqBodyIdentityCenterStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetupLandingZoneReqBodyIdentityCenterStatus) UnmarshalJSON(b []byte) error {
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

type SetupLandingZoneReqBodyOrganizationStructureType struct {
	value string
}

type SetupLandingZoneReqBodyOrganizationStructureTypeEnum struct {
	STANDARD     SetupLandingZoneReqBodyOrganizationStructureType
	NON_STANDARD SetupLandingZoneReqBodyOrganizationStructureType
}

func GetSetupLandingZoneReqBodyOrganizationStructureTypeEnum() SetupLandingZoneReqBodyOrganizationStructureTypeEnum {
	return SetupLandingZoneReqBodyOrganizationStructureTypeEnum{
		STANDARD: SetupLandingZoneReqBodyOrganizationStructureType{
			value: "STANDARD",
		},
		NON_STANDARD: SetupLandingZoneReqBodyOrganizationStructureType{
			value: "NON_STANDARD",
		},
	}
}

func (c SetupLandingZoneReqBodyOrganizationStructureType) Value() string {
	return c.value
}

func (c SetupLandingZoneReqBodyOrganizationStructureType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetupLandingZoneReqBodyOrganizationStructureType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLatestUpgradableReleaseRequest Request Object
type ShowLatestUpgradableReleaseRequest struct {

	// 子产品名称
	SubProductName string `json:"sub_product_name"`

	// 系统类型
	OsType ShowLatestUpgradableReleaseRequestOsType `json:"os_type"`

	// CPU架构
	Arch *ShowLatestUpgradableReleaseRequestArch `json:"arch,omitempty"`
}

func (o ShowLatestUpgradableReleaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestUpgradableReleaseRequest struct{}"
	}

	return strings.Join([]string{"ShowLatestUpgradableReleaseRequest", string(data)}, " ")
}

type ShowLatestUpgradableReleaseRequestOsType struct {
	value string
}

type ShowLatestUpgradableReleaseRequestOsTypeEnum struct {
	WINDOWS ShowLatestUpgradableReleaseRequestOsType
	DEBIAN  ShowLatestUpgradableReleaseRequestOsType
	DARWIN  ShowLatestUpgradableReleaseRequestOsType
}

func GetShowLatestUpgradableReleaseRequestOsTypeEnum() ShowLatestUpgradableReleaseRequestOsTypeEnum {
	return ShowLatestUpgradableReleaseRequestOsTypeEnum{
		WINDOWS: ShowLatestUpgradableReleaseRequestOsType{
			value: "windows",
		},
		DEBIAN: ShowLatestUpgradableReleaseRequestOsType{
			value: "debian",
		},
		DARWIN: ShowLatestUpgradableReleaseRequestOsType{
			value: "darwin",
		},
	}
}

func (c ShowLatestUpgradableReleaseRequestOsType) Value() string {
	return c.value
}

func (c ShowLatestUpgradableReleaseRequestOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLatestUpgradableReleaseRequestOsType) UnmarshalJSON(b []byte) error {
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

type ShowLatestUpgradableReleaseRequestArch struct {
	value string
}

type ShowLatestUpgradableReleaseRequestArchEnum struct {
	X86_64    ShowLatestUpgradableReleaseRequestArch
	ARM64     ShowLatestUpgradableReleaseRequestArch
	UNIVERSAL ShowLatestUpgradableReleaseRequestArch
}

func GetShowLatestUpgradableReleaseRequestArchEnum() ShowLatestUpgradableReleaseRequestArchEnum {
	return ShowLatestUpgradableReleaseRequestArchEnum{
		X86_64: ShowLatestUpgradableReleaseRequestArch{
			value: "X86-64",
		},
		ARM64: ShowLatestUpgradableReleaseRequestArch{
			value: "Arm64",
		},
		UNIVERSAL: ShowLatestUpgradableReleaseRequestArch{
			value: "Universal",
		},
	}
}

func (c ShowLatestUpgradableReleaseRequestArch) Value() string {
	return c.value
}

func (c ShowLatestUpgradableReleaseRequestArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLatestUpgradableReleaseRequestArch) UnmarshalJSON(b []byte) error {
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

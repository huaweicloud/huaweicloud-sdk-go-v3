package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VolumeFlavorItemDetail struct {

	// **参数解释**: IO类型。 **取值范围**: - ULTRAHIGH：超高IO。 - ESSD：极速型SSD。
	IoType *VolumeFlavorItemDetailIoType `json:"io_type,omitempty"`

	// **参数解释**: 存储介质类型。 **取值范围**: - SAS：高IO云硬盘。 - SSD：超高IO云硬盘。 - ESSD：极速型SSD云硬盘。
	MediaType *VolumeFlavorItemDetailMediaType `json:"media_type,omitempty"`

	// **参数解释**: 存储规格在各个可用区的状态。 **取值范围**: - normal：正常。 - sellout：已售罄。 - unknown：状态未知。 - unsupported：不支持。 - abandon：已废弃。
	AzStatus map[string]string `json:"az_status,omitempty"`

	// **参数解释**: 该存储规格支持的计算规格组类型列表。 **取值范围**: 不涉及。
	SupportComputeGroupType *[]string `json:"support_compute_group_type,omitempty"`
}

func (o VolumeFlavorItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeFlavorItemDetail struct{}"
	}

	return strings.Join([]string{"VolumeFlavorItemDetail", string(data)}, " ")
}

type VolumeFlavorItemDetailIoType struct {
	value string
}

type VolumeFlavorItemDetailIoTypeEnum struct {
	ULTRAHIGH VolumeFlavorItemDetailIoType
	ESSD      VolumeFlavorItemDetailIoType
}

func GetVolumeFlavorItemDetailIoTypeEnum() VolumeFlavorItemDetailIoTypeEnum {
	return VolumeFlavorItemDetailIoTypeEnum{
		ULTRAHIGH: VolumeFlavorItemDetailIoType{
			value: "ULTRAHIGH",
		},
		ESSD: VolumeFlavorItemDetailIoType{
			value: "ESSD",
		},
	}
}

func (c VolumeFlavorItemDetailIoType) Value() string {
	return c.value
}

func (c VolumeFlavorItemDetailIoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeFlavorItemDetailIoType) UnmarshalJSON(b []byte) error {
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

type VolumeFlavorItemDetailMediaType struct {
	value string
}

type VolumeFlavorItemDetailMediaTypeEnum struct {
	SAS  VolumeFlavorItemDetailMediaType
	SSD  VolumeFlavorItemDetailMediaType
	ESSD VolumeFlavorItemDetailMediaType
}

func GetVolumeFlavorItemDetailMediaTypeEnum() VolumeFlavorItemDetailMediaTypeEnum {
	return VolumeFlavorItemDetailMediaTypeEnum{
		SAS: VolumeFlavorItemDetailMediaType{
			value: "SAS",
		},
		SSD: VolumeFlavorItemDetailMediaType{
			value: "SSD",
		},
		ESSD: VolumeFlavorItemDetailMediaType{
			value: "ESSD",
		},
	}
}

func (c VolumeFlavorItemDetailMediaType) Value() string {
	return c.value
}

func (c VolumeFlavorItemDetailMediaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeFlavorItemDetailMediaType) UnmarshalJSON(b []byte) error {
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

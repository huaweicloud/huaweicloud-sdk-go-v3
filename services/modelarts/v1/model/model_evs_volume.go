package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EvsVolume struct {

	// **参数解释**：系统盘大小。表示分配给系统盘的存储空间大小。 **约束限制**：不涉及。 **取值范围**：100 - 1024 GB **默认取值**：不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**：系统盘类型。表示系统盘的存储类型。 **约束限制**：不涉及。 **取值范围**： - ESSD：极速型SSD云硬盘 - GPSSD：通用型SSD云硬盘 - SAS：高IO云硬盘 - SATA：普通IO云硬盘 - SSD：超高IO云硬盘 **默认取值**：不涉及。
	Type *EvsVolumeType `json:"type,omitempty"`
}

func (o EvsVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvsVolume struct{}"
	}

	return strings.Join([]string{"EvsVolume", string(data)}, " ")
}

type EvsVolumeType struct {
	value string
}

type EvsVolumeTypeEnum struct {
	ESSD  EvsVolumeType
	GPSSD EvsVolumeType
	SAS   EvsVolumeType
	SATA  EvsVolumeType
	SSD   EvsVolumeType
}

func GetEvsVolumeTypeEnum() EvsVolumeTypeEnum {
	return EvsVolumeTypeEnum{
		ESSD: EvsVolumeType{
			value: "ESSD",
		},
		GPSSD: EvsVolumeType{
			value: "GPSSD",
		},
		SAS: EvsVolumeType{
			value: "SAS",
		},
		SATA: EvsVolumeType{
			value: "SATA",
		},
		SSD: EvsVolumeType{
			value: "SSD",
		},
	}
}

func (c EvsVolumeType) Value() string {
	return c.value
}

func (c EvsVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EvsVolumeType) UnmarshalJSON(b []byte) error {
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

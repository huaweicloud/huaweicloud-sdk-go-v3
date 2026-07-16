package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ServerVolume struct {

	// **参数解释**：EVS盘大小。表示分配给系统盘的存储空间大小。 **约束限制**：不涉及。 **取值范围**：100 - 1024 GB **默认取值**：不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**：存储类型。表示系统盘或数据盘。 **约束限制**：不涉及。 **取值范围**： - ROOT：系统盘 - DATA：数据盘  **默认取值**：不涉及。
	Type *ServerVolumeType `json:"type,omitempty"`

	// **参数解释**：EVS盘类型。表示EVS盘的存储类型。 **约束限制**：不涉及。 **取值范围**： - ESSD：极速型SSD云硬盘 - GPSSD：通用型SSD云硬盘 - SAS：高IO云硬盘 - SATA：普通IO云硬盘 - SSD：超高IO云硬盘  **默认取值**：不涉及。
	EvsType *ServerVolumeEvsType `json:"evs_type,omitempty"`

	// **参数解释**：EVS盘的ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	EvsId *string `json:"evs_id,omitempty"`
}

func (o ServerVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerVolume struct{}"
	}

	return strings.Join([]string{"ServerVolume", string(data)}, " ")
}

type ServerVolumeType struct {
	value string
}

type ServerVolumeTypeEnum struct {
	ROOT ServerVolumeType
	DATA ServerVolumeType
}

func GetServerVolumeTypeEnum() ServerVolumeTypeEnum {
	return ServerVolumeTypeEnum{
		ROOT: ServerVolumeType{
			value: "ROOT",
		},
		DATA: ServerVolumeType{
			value: "DATA",
		},
	}
}

func (c ServerVolumeType) Value() string {
	return c.value
}

func (c ServerVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerVolumeType) UnmarshalJSON(b []byte) error {
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

type ServerVolumeEvsType struct {
	value string
}

type ServerVolumeEvsTypeEnum struct {
	ESSD  ServerVolumeEvsType
	GPSSD ServerVolumeEvsType
	SAS   ServerVolumeEvsType
	SATA  ServerVolumeEvsType
	SSD   ServerVolumeEvsType
}

func GetServerVolumeEvsTypeEnum() ServerVolumeEvsTypeEnum {
	return ServerVolumeEvsTypeEnum{
		ESSD: ServerVolumeEvsType{
			value: "ESSD",
		},
		GPSSD: ServerVolumeEvsType{
			value: "GPSSD",
		},
		SAS: ServerVolumeEvsType{
			value: "SAS",
		},
		SATA: ServerVolumeEvsType{
			value: "SATA",
		},
		SSD: ServerVolumeEvsType{
			value: "SSD",
		},
	}
}

func (c ServerVolumeEvsType) Value() string {
	return c.value
}

func (c ServerVolumeEvsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerVolumeEvsType) UnmarshalJSON(b []byte) error {
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

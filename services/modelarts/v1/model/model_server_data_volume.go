package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ServerDataVolume 创建Lite Server的数据盘信息。
type ServerDataVolume struct {

	// **参数解释**：数据盘大小。表示分配给数据盘的存储空间大小。 **约束限制**：不涉及。 **取值范围**：100 - 32768 GB **默认取值**：不涉及。
	Size int32 `json:"size"`

	// **参数解释**：系统盘类型。表示数据盘的存储类型。 **约束限制**：不涉及。 **取值范围**： - ESSD：极速型SSD云硬盘 - GPSSD：通用型SSD云硬盘 - SAS：高IO云硬盘 - SATA：普通IO云硬盘 - SSD：超高IO云硬盘 **默认取值**：不涉及。
	Type ServerDataVolumeType `json:"type"`

	// **参数解释**：数据盘个数。表示为实例分配的数据盘数量。 **约束限制**：不涉及。 **取值范围**：1 - 8 **默认取值**：不涉及。
	Count int32 `json:"count"`
}

func (o ServerDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerDataVolume struct{}"
	}

	return strings.Join([]string{"ServerDataVolume", string(data)}, " ")
}

type ServerDataVolumeType struct {
	value string
}

type ServerDataVolumeTypeEnum struct {
	SSD   ServerDataVolumeType
	SAS   ServerDataVolumeType
	SATA  ServerDataVolumeType
	GPSSD ServerDataVolumeType
	ESSD  ServerDataVolumeType
}

func GetServerDataVolumeTypeEnum() ServerDataVolumeTypeEnum {
	return ServerDataVolumeTypeEnum{
		SSD: ServerDataVolumeType{
			value: "SSD",
		},
		SAS: ServerDataVolumeType{
			value: "SAS",
		},
		SATA: ServerDataVolumeType{
			value: "SATA",
		},
		GPSSD: ServerDataVolumeType{
			value: "GPSSD",
		},
		ESSD: ServerDataVolumeType{
			value: "ESSD",
		},
	}
}

func (c ServerDataVolumeType) Value() string {
	return c.value
}

func (c ServerDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerDataVolumeType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VolumeType 磁盘类型，获取可用磁盘类型详见接口磁盘管理ListVolumeType。 * `ESSD` - 极速型SSD * `SSD` - 超高IO * `GPSSD` - 通用型SSD * `SAS` - 高IO * `SATA` - 普通IO
type VolumeType struct {
	value string
}

type VolumeTypeEnum struct {
	ESSD  VolumeType
	SSD   VolumeType
	GPSSD VolumeType
	SAS   VolumeType
	SATA  VolumeType
}

func GetVolumeTypeEnum() VolumeTypeEnum {
	return VolumeTypeEnum{
		ESSD: VolumeType{
			value: "ESSD",
		},
		SSD: VolumeType{
			value: "SSD",
		},
		GPSSD: VolumeType{
			value: "GPSSD",
		},
		SAS: VolumeType{
			value: "SAS",
		},
		SATA: VolumeType{
			value: "SATA",
		},
	}
}

func (c VolumeType) Value() string {
	return c.value
}

func (c VolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeType) UnmarshalJSON(b []byte) error {
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

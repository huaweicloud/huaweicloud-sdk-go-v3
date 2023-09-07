package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VolumeKindObj API类型，固定值“Volume”，该值不可修改。
type VolumeKindObj struct {
	value string
}

type VolumeKindObjEnum struct {
	VOLUME VolumeKindObj
}

func GetVolumeKindObjEnum() VolumeKindObjEnum {
	return VolumeKindObjEnum{
		VOLUME: VolumeKindObj{
			value: "Volume",
		},
	}
}

func (c VolumeKindObj) Value() string {
	return c.value
}

func (c VolumeKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeKindObj) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SupabaseVolumeConfig 存储配置。
type SupabaseVolumeConfig struct {

	// 磁盘类型。
	Type *SupabaseVolumeConfigType `json:"type,omitempty"`

	// 磁盘大小（GB）。
	Size *int32 `json:"size,omitempty"`
}

func (o SupabaseVolumeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseVolumeConfig struct{}"
	}

	return strings.Join([]string{"SupabaseVolumeConfig", string(data)}, " ")
}

type SupabaseVolumeConfigType struct {
	value string
}

type SupabaseVolumeConfigTypeEnum struct {
	CLOUDSSD SupabaseVolumeConfigType
}

func GetSupabaseVolumeConfigTypeEnum() SupabaseVolumeConfigTypeEnum {
	return SupabaseVolumeConfigTypeEnum{
		CLOUDSSD: SupabaseVolumeConfigType{
			value: "CLOUDSSD",
		},
	}
}

func (c SupabaseVolumeConfigType) Value() string {
	return c.value
}

func (c SupabaseVolumeConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SupabaseVolumeConfigType) UnmarshalJSON(b []byte) error {
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

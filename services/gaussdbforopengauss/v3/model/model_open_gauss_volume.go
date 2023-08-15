package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussVolume volume信息。
type OpenGaussVolume struct {

	// 磁盘类型。  仅支持ULTRAHIGH和ESSD，区分大小写，分别表示SSD和急速云盘。
	Type OpenGaussVolumeType `json:"type"`

	// 磁盘大小。例如：该参数填写为“40”，表示为创建的实例分配40GB的磁盘空间。  取值范围：（分片数*40GB）~（分片数*16TB），且大小只能为分片数*40的整数倍。
	Size int32 `json:"size"`
}

func (o OpenGaussVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussVolume struct{}"
	}

	return strings.Join([]string{"OpenGaussVolume", string(data)}, " ")
}

type OpenGaussVolumeType struct {
	value string
}

type OpenGaussVolumeTypeEnum struct {
	ULTRAHIGH OpenGaussVolumeType
	ESSD      OpenGaussVolumeType
}

func GetOpenGaussVolumeTypeEnum() OpenGaussVolumeTypeEnum {
	return OpenGaussVolumeTypeEnum{
		ULTRAHIGH: OpenGaussVolumeType{
			value: "ULTRAHIGH",
		},
		ESSD: OpenGaussVolumeType{
			value: "ESSD",
		},
	}
}

func (c OpenGaussVolumeType) Value() string {
	return c.value
}

func (c OpenGaussVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussVolumeType) UnmarshalJSON(b []byte) error {
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

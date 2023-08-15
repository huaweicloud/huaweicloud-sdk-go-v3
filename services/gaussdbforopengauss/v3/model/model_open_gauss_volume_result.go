package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussVolumeResult volume信息。
type OpenGaussVolumeResult struct {

	// 磁盘类型。  取值如下，区分大小写：  - ULTRAHIGH，表示SSD。 - ESSD，表示急速云盘
	Type OpenGaussVolumeResultType `json:"type"`

	// 磁盘大小。  GaussDB分布式实例创建时需指定大小：要求必须为（分片数 * 40GB）的倍数，取值范围：（分片数*40GB）~（分片数*16TB）。
	Size int32 `json:"size"`
}

func (o OpenGaussVolumeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussVolumeResult struct{}"
	}

	return strings.Join([]string{"OpenGaussVolumeResult", string(data)}, " ")
}

type OpenGaussVolumeResultType struct {
	value string
}

type OpenGaussVolumeResultTypeEnum struct {
	ULTRAHIGH OpenGaussVolumeResultType
	ESSD      OpenGaussVolumeResultType
}

func GetOpenGaussVolumeResultTypeEnum() OpenGaussVolumeResultTypeEnum {
	return OpenGaussVolumeResultTypeEnum{
		ULTRAHIGH: OpenGaussVolumeResultType{
			value: "ULTRAHIGH",
		},
		ESSD: OpenGaussVolumeResultType{
			value: "ESSD",
		},
	}
}

func (c OpenGaussVolumeResultType) Value() string {
	return c.value
}

func (c OpenGaussVolumeResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussVolumeResultType) UnmarshalJSON(b []byte) error {
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

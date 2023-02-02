package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// volume信息。
type OpenGaussVolumeResponse struct {

	// 磁盘类型。  取值如下，区分大小写：  - ULTRAHIGH，表示SSD。 - ESSD，表示急速云盘
	Type OpenGaussVolumeResponseType `json:"type"`

	// 磁盘大小。  GaussDB分布式实例创建时需指定大小：要求必须为（分片数 * 40GB）的倍数，取值范围：（分片数*40GB）~（分片数*16TB）。
	Size int32 `json:"size"`
}

func (o OpenGaussVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussVolumeResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussVolumeResponse", string(data)}, " ")
}

type OpenGaussVolumeResponseType struct {
	value string
}

type OpenGaussVolumeResponseTypeEnum struct {
	ULTRAHIGH OpenGaussVolumeResponseType
	ESSD      OpenGaussVolumeResponseType
}

func GetOpenGaussVolumeResponseTypeEnum() OpenGaussVolumeResponseTypeEnum {
	return OpenGaussVolumeResponseTypeEnum{
		ULTRAHIGH: OpenGaussVolumeResponseType{
			value: "ULTRAHIGH",
		},
		ESSD: OpenGaussVolumeResponseType{
			value: "ESSD",
		},
	}
}

func (c OpenGaussVolumeResponseType) Value() string {
	return c.value
}

func (c OpenGaussVolumeResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussVolumeResponseType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

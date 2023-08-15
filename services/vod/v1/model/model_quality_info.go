package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QualityInfo
type QualityInfo struct {
	Video *VideoTemplateInfo `json:"video,omitempty"`

	Audio *AudioTemplateInfo `json:"audio,omitempty"`

	// 格式。
	Format QualityInfoFormat `json:"format"`
}

func (o QualityInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualityInfo struct{}"
	}

	return strings.Join([]string{"QualityInfo", string(data)}, " ")
}

type QualityInfoFormat struct {
	value string
}

type QualityInfoFormatEnum struct {
	MP4      QualityInfoFormat
	HLS      QualityInfoFormat
	DASH     QualityInfoFormat
	DASH_HLS QualityInfoFormat
	MP3      QualityInfoFormat
	ADTS     QualityInfoFormat
	UNKNOW   QualityInfoFormat
}

func GetQualityInfoFormatEnum() QualityInfoFormatEnum {
	return QualityInfoFormatEnum{
		MP4: QualityInfoFormat{
			value: "MP4",
		},
		HLS: QualityInfoFormat{
			value: "HLS",
		},
		DASH: QualityInfoFormat{
			value: "DASH",
		},
		DASH_HLS: QualityInfoFormat{
			value: "DASH_HLS",
		},
		MP3: QualityInfoFormat{
			value: "MP3",
		},
		ADTS: QualityInfoFormat{
			value: "ADTS",
		},
		UNKNOW: QualityInfoFormat{
			value: "UNKNOW",
		},
	}
}

func (c QualityInfoFormat) Value() string {
	return c.value
}

func (c QualityInfoFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QualityInfoFormat) UnmarshalJSON(b []byte) error {
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

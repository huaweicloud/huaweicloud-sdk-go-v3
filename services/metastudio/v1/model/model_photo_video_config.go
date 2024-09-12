package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PhotoVideoConfig 视频输出配置。照片数字人支持该输出配置查询，不支持修改。
type PhotoVideoConfig struct {

	// 视频编码格式及视频文件格式。 * H264：h264编码，输出mp4文件
	Codec PhotoVideoConfigCodec `json:"codec"`

	// **参数解释**： 输出平均码率。  单位：kbps。  最小值40，最大值30000。
	Bitrate *int32 `json:"bitrate,omitempty"`

	// 帧率。  单位：FPS。
	FrameRate *PhotoVideoConfigFrameRate `json:"frame_rate,omitempty"`
}

func (o PhotoVideoConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhotoVideoConfig struct{}"
	}

	return strings.Join([]string{"PhotoVideoConfig", string(data)}, " ")
}

type PhotoVideoConfigCodec struct {
	value string
}

type PhotoVideoConfigCodecEnum struct {
	H264 PhotoVideoConfigCodec
}

func GetPhotoVideoConfigCodecEnum() PhotoVideoConfigCodecEnum {
	return PhotoVideoConfigCodecEnum{
		H264: PhotoVideoConfigCodec{
			value: "H264",
		},
	}
}

func (c PhotoVideoConfigCodec) Value() string {
	return c.value
}

func (c PhotoVideoConfigCodec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PhotoVideoConfigCodec) UnmarshalJSON(b []byte) error {
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

type PhotoVideoConfigFrameRate struct {
	value string
}

type PhotoVideoConfigFrameRateEnum struct {
	E_24 PhotoVideoConfigFrameRate
	E_25 PhotoVideoConfigFrameRate
	E_30 PhotoVideoConfigFrameRate
}

func GetPhotoVideoConfigFrameRateEnum() PhotoVideoConfigFrameRateEnum {
	return PhotoVideoConfigFrameRateEnum{
		E_24: PhotoVideoConfigFrameRate{
			value: "24",
		},
		E_25: PhotoVideoConfigFrameRate{
			value: "25",
		},
		E_30: PhotoVideoConfigFrameRate{
			value: "30",
		},
	}
}

func (c PhotoVideoConfigFrameRate) Value() string {
	return c.value
}

func (c PhotoVideoConfigFrameRate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PhotoVideoConfigFrameRate) UnmarshalJSON(b []byte) error {
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

package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Multidrm struct {
	// 唯一标识

	ContentId *string `json:"content_id,omitempty"`
	// 定义数据流的类型，取值为DASH或HLS

	StreamingMode MultidrmStreamingMode `json:"streaming_mode"`
	// 音频加密开关。  取值如下： - 0：标识音频不加密。 - 1：标识音频加密。  默认值：0。 该参数只对dash有效。

	EncryptAudio *int32 `json:"encrypt_audio,omitempty"`
	// 定义加密方式。  取值如下： - 16418(AES-128,CBC) - 16420(AES-128,CTR) - 16422(SM4CBC)

	Emi *int32 `json:"emi,omitempty"`
	// HLS视频加密控制参数。  取值如下： - PLAYREADY - CHINA_DRM - WIDEVINE

	DrmList []string `json:"drm_list"`
}

func (o Multidrm) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Multidrm struct{}"
	}

	return strings.Join([]string{"Multidrm", string(data)}, " ")
}

type MultidrmStreamingMode struct {
	value string
}

type MultidrmStreamingModeEnum struct {
	DASH MultidrmStreamingMode
	HLS  MultidrmStreamingMode
}

func GetMultidrmStreamingModeEnum() MultidrmStreamingModeEnum {
	return MultidrmStreamingModeEnum{
		DASH: MultidrmStreamingMode{
			value: "DASH",
		},
		HLS: MultidrmStreamingMode{
			value: "HLS",
		},
	}
}

func (c MultidrmStreamingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *MultidrmStreamingMode) UnmarshalJSON(b []byte) error {
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

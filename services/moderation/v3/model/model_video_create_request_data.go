package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VideoCreateRequestData 视频数据输入
type VideoCreateRequestData struct {

	// 视频url地址
	Url string `json:"url"`

	// 截帧频率间隔，单位为秒，取值范围为1~60s；若不传递默认5s截帧一次
	FrameInterval *int32 `json:"frame_interval,omitempty"`

	// 支持的语言，默认为zh zh：中文
	Language *VideoCreateRequestDataLanguage `json:"language,omitempty"`
}

func (o VideoCreateRequestData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCreateRequestData struct{}"
	}

	return strings.Join([]string{"VideoCreateRequestData", string(data)}, " ")
}

type VideoCreateRequestDataLanguage struct {
	value string
}

type VideoCreateRequestDataLanguageEnum struct {
	ZH VideoCreateRequestDataLanguage
}

func GetVideoCreateRequestDataLanguageEnum() VideoCreateRequestDataLanguageEnum {
	return VideoCreateRequestDataLanguageEnum{
		ZH: VideoCreateRequestDataLanguage{
			value: "zh",
		},
	}
}

func (c VideoCreateRequestDataLanguage) Value() string {
	return c.value
}

func (c VideoCreateRequestDataLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoCreateRequestDataLanguage) UnmarshalJSON(b []byte) error {
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

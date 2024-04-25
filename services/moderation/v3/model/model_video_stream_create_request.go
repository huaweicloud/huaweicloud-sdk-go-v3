package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoStreamCreateRequest struct {
	Data *VideoStreamCreateRequestData `json:"data"`

	// 事件类型，可选值如下： default：默认事件
	EventType VideoStreamCreateRequestEventType `json:"event_type"`

	// 视频流中画面需要检测的风险类型，列表不能为空。 porn：鉴黄内容的检测 terrorism：涉政暴恐内容的检测 politics：政治敏感人物内容的检测 image_text：图文违规内容的检测（检测图片中出现的广告、色情、暴恐、涉政的文字违规内容以及二维码内容）
	ImageCategories []VideoStreamCreateRequestImageCategories `json:"image_categories"`

	// 视频流中音频需要检测的风险类型，不传或为空时表示不审核音频维度。 porn：涉黄检测 politics: 涉政检测 abuse: 辱骂检测 ad: 广告检测 moan: 娇喘检测
	AudioCategories *[]VideoStreamCreateRequestAudioCategories `json:"audio_categories,omitempty"`

	// 回调http接口，回调内容如下： ```{     \"job_id\":\"xxxxxx\",     \"status\":\"running\", //running - 审核中，succeeded - 审核完成，failed - 审核失败     \"request_id\":\"2419446b1fe14203f64e4018d12db3dd\",     \"craete_time\":\"2022-07-30T08:57:11.011Z\",     \"update_time\":\"2022-07-30T08:57:14.014Z\",     \"result\":{         \"suggestion\":\"block\",         \"audio_detail\":[             {                 \"suggestion\":\"block\",                 \"label\":\"politics\",                 \"audio_text\":\"xxxx\",                 \"detail\":[                     {                         \"confidence\":1,                         \"suggestion\":\"block\",                         \"label\":\"politics\",                         \"segments\":[                             {                                 \"segment\":\"xxx\"                             },                             {                                 \"segment\":\"xxx\"                             },                             {                                 \"segment\":\"xxx\"                             }                         ]                     }                 ]             }         ],         \"image_detail\":[             {                 \"suggestion\":\"block\",                 \"category\":\"xxx\",                 \"time\":\"xxx\", // 视频流截帧图片发生的绝对时间                 \"detail\":[                     {                         \"face_location\":{                             \"bottom_right_x\":247,                             \"bottom_right_y\":191,                             \"top_left_y\":79,                             \"top_left_x\":160                         },                         \"confidence\":1,                         \"suggestion\":\"block\",                         \"label\":\"xxx\",                         \"category\":\"xxx\"                     },                     {                         \"qr_content\":\"xxx\",                         \"confidence\":\"xxx\",                         \"suggestion\":\"xxx\",                         \"label\":\"xxx\",                         \"category\":\"xxx\",                         \"qr_location\":{                             \"bottom_right_x\":554,                             \"bottom_right_y\":842,                             \"top_left_y\":426,                             \"top_left_x\":140                         }                     },                     {                         \"confidence\":1,                         \"suggestion\":\"block\",                         \"label\":\"politics\",                         \"category\":\"image_text\",                         \"segments\":[                             {                                 \"segment\":\"x\"                             },                             {                                 \"segment\":\"xxx\"                             },                             {                                 \"segment\":\"xx\"                             },                             {                                 \"segment\":\"x\"                             },                             {                                 \"segment\":\"xxxx\"                             }                         ]                     }                 ]             }         ]     },     \"request_params\":{         \"data\":{             \"url\":\"rtmp://xxxx\",             \"frame_interval\":3         },         \"event_type\":\"default\",         \"image_categories\":[             \"politics\",             \"porn\",             \"image_text\",             \"terrorism\"         ],         \"audio_categories\":[             \"porn\",             \"ad\",             \"politics\",             \"moan\",             \"abuse\"         ],         \"callback\":\"http://xxx/callback\"     } }
	Callback string `json:"callback"`

	// 用于回调通知时校验请求由华为云内容安全服务发起，由您自定义。随机字符串，由英文字母、数字、下划线组成，不超过64个字符。 说明：当seed非空时，headers中将包含X-Auth-Signature字段，字段的值使用HmacSHA256算法生成，待加密字符串由create_time、job_id、request_id、seed按照顺序拼接而成，密钥为seed。
	Seed *string `json:"seed,omitempty"`
}

func (o VideoStreamCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoStreamCreateRequest struct{}"
	}

	return strings.Join([]string{"VideoStreamCreateRequest", string(data)}, " ")
}

type VideoStreamCreateRequestEventType struct {
	value string
}

type VideoStreamCreateRequestEventTypeEnum struct {
	DEFAULT VideoStreamCreateRequestEventType
}

func GetVideoStreamCreateRequestEventTypeEnum() VideoStreamCreateRequestEventTypeEnum {
	return VideoStreamCreateRequestEventTypeEnum{
		DEFAULT: VideoStreamCreateRequestEventType{
			value: "default",
		},
	}
}

func (c VideoStreamCreateRequestEventType) Value() string {
	return c.value
}

func (c VideoStreamCreateRequestEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoStreamCreateRequestEventType) UnmarshalJSON(b []byte) error {
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

type VideoStreamCreateRequestImageCategories struct {
	value string
}

type VideoStreamCreateRequestImageCategoriesEnum struct {
	PORN       VideoStreamCreateRequestImageCategories
	POLITICS   VideoStreamCreateRequestImageCategories
	TERRORISM  VideoStreamCreateRequestImageCategories
	IMAGE_TEXT VideoStreamCreateRequestImageCategories
}

func GetVideoStreamCreateRequestImageCategoriesEnum() VideoStreamCreateRequestImageCategoriesEnum {
	return VideoStreamCreateRequestImageCategoriesEnum{
		PORN: VideoStreamCreateRequestImageCategories{
			value: "porn",
		},
		POLITICS: VideoStreamCreateRequestImageCategories{
			value: "politics",
		},
		TERRORISM: VideoStreamCreateRequestImageCategories{
			value: "terrorism",
		},
		IMAGE_TEXT: VideoStreamCreateRequestImageCategories{
			value: "image_text",
		},
	}
}

func (c VideoStreamCreateRequestImageCategories) Value() string {
	return c.value
}

func (c VideoStreamCreateRequestImageCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoStreamCreateRequestImageCategories) UnmarshalJSON(b []byte) error {
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

type VideoStreamCreateRequestAudioCategories struct {
	value string
}

type VideoStreamCreateRequestAudioCategoriesEnum struct {
	PORN     VideoStreamCreateRequestAudioCategories
	AD       VideoStreamCreateRequestAudioCategories
	POLITICS VideoStreamCreateRequestAudioCategories
	MOAN     VideoStreamCreateRequestAudioCategories
	ABUSE    VideoStreamCreateRequestAudioCategories
}

func GetVideoStreamCreateRequestAudioCategoriesEnum() VideoStreamCreateRequestAudioCategoriesEnum {
	return VideoStreamCreateRequestAudioCategoriesEnum{
		PORN: VideoStreamCreateRequestAudioCategories{
			value: "porn",
		},
		AD: VideoStreamCreateRequestAudioCategories{
			value: "ad",
		},
		POLITICS: VideoStreamCreateRequestAudioCategories{
			value: "politics",
		},
		MOAN: VideoStreamCreateRequestAudioCategories{
			value: "moan",
		},
		ABUSE: VideoStreamCreateRequestAudioCategories{
			value: "abuse",
		},
	}
}

func (c VideoStreamCreateRequestAudioCategories) Value() string {
	return c.value
}

func (c VideoStreamCreateRequestAudioCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoStreamCreateRequestAudioCategories) UnmarshalJSON(b []byte) error {
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

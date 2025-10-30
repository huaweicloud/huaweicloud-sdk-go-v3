package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWatermarkTemplateResponse Response Object
type ShowWatermarkTemplateResponse struct {

	// 水印模板名称
	Name *string `json:"name,omitempty"`

	// 水印类型，0图片，1文字,2图文共存
	Type *int32 `json:"type,omitempty"`

	// 模板注释
	Description *string `json:"description,omitempty"`

	// 图片下载路径 校验规则：图片URL长度大于0，最大长度2048，能够正常按URL格式解析，必须是 http 或 https 协议。（OTT 场景限制只支持https） 图片格式: .png/.jpg/.PNG/.JPG结尾
	PictureUrl *string `json:"picture_url,omitempty"`

	// 1）整数型代表水印图片宽的像素值，范围0或[8，4096]，单位px。 2）小数型代表相对输出视频分辨率宽的比率，范围[0,1) 建议宽高只设置一项，另外一项会自适应缩放，避免变形。宽高全部设置为0表示水印图片原始宽高 百分比限制最多保留小数点后4位，width和height 只支持同时为像素或是百分比，不支持一个像素，一个百分比
	Width float32 `json:"width,omitempty"`

	// 水印图片高，值有两种形式： 1）整数型代表水印图片高的像素值，范围0或[8，4096]，单位px。 2）小数型代表相对输出视频分辨率高的比率，范围[0，1) 建议宽高只设置一项，另外一项会自适应缩放，避免变形。宽高全部设置为0表示水印图片原始宽高 百分比限制最多保留小数点后4位，width和height 只支持同时为像素或是百分比，不支持一个像素，一个百分比
	Height float32 `json:"height,omitempty"`

	Location *WatermarkLocation `json:"location,omitempty"`

	Text *WordWaterMarkInfo `json:"text,omitempty"`

	// 业务属性，cloud_live：云直播，默认值；media_live：媒体直播，不支持修改
	Scene *ShowWatermarkTemplateResponseScene `json:"scene,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWatermarkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowWatermarkTemplateResponse", string(data)}, " ")
}

type ShowWatermarkTemplateResponseScene struct {
	value string
}

type ShowWatermarkTemplateResponseSceneEnum struct {
	CLOUD_LIVE ShowWatermarkTemplateResponseScene
	MEDIA_LIVE ShowWatermarkTemplateResponseScene
}

func GetShowWatermarkTemplateResponseSceneEnum() ShowWatermarkTemplateResponseSceneEnum {
	return ShowWatermarkTemplateResponseSceneEnum{
		CLOUD_LIVE: ShowWatermarkTemplateResponseScene{
			value: "cloud_live",
		},
		MEDIA_LIVE: ShowWatermarkTemplateResponseScene{
			value: "media_live",
		},
	}
}

func (c ShowWatermarkTemplateResponseScene) Value() string {
	return c.value
}

func (c ShowWatermarkTemplateResponseScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWatermarkTemplateResponseScene) UnmarshalJSON(b []byte) error {
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

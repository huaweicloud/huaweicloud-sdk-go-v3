package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 合流参数  - 纯音频录制  encode_template 填 audio_only，音频合流会动态选择最大三方的声音。  layout_template、layout_panes 以及其他视频相关参数都不填，填就忽略。  - 音视频录制（包括共享桌面）  encode_template 非 audio_only，layout_template必须非空。  音频合流会动态选择最大三方的声音
type MixParam struct {
	// 房间id

	RoomId string `json:"room_id"`
	// 输出编码模板名称 - 1920x1080_30_4620：输出流是1080p，帧率30，码率4.62Mbps - 1920x1080_15_3460：输出流是1080p，帧率15，码率3.46Mbps - 1280x720_30_3420：输出流是720p，帧率30，码率3.42Mbps - 1280x720_15_2260：输出流是720p，帧率15，码率2.26Mbps

	EncodeTemplate MixParamEncodeTemplate `json:"encode_template"`
	// 最长空闲频道时间。  取值范围：[5，43200]，默认值为30。  单位：秒。  如果频道内无连麦方的状态持续超过该时间，录制程序会自动退出。退出后，再次调用start请求，会产生新的录制任务。  连麦方指：joiner或者publisher的用户。

	MaxIdleTime *int32 `json:"max_idle_time,omitempty"`
	// 视频布局模板编号，不混视频的时候，不需要带。 - nine_grids_view：九宫格模板（自适应模板） - screen_share_left：主视图在左边的屏幕共享模板（自适应模板） - screen_share_right：主视图在右边的屏幕共享模板（自适应模板） - custom：自定义布局

	LayoutTemplate *string `json:"layout_template,omitempty"`
	// 默认用户背景图地址，图片先上传obs，格式s3://bucket/object。默认使用等比缩放裁剪，保证铺满。

	DefaultUserBackgroundImage *string `json:"default_user_background_image,omitempty"`
	// 共享屏幕的背景图地址，图片先上传obs，格式s3://bucket/object。  在一大多小的布局场景下，无论大窗是显示非指定用户（屏幕共享人的桌面）还是指定用户的共享桌面，都通过该字段指定背景图。

	ScreenBackgroundImage *string `json:"screen_background_image,omitempty"`
	// 画布背景图地址，图片先上传obs，格式s3://bucket/object。默认使用等比缩放裁剪，保证铺满。

	BackgroundImage *string `json:"background_image,omitempty"`
	// 需要混流的视频列表。若不需要混流视频，则可不传递该参数。  nine_grids_view模板不需要填写本字段。

	LayoutPanes *[]MixLayoutPane `json:"layout_panes,omitempty"`
	// 指定用户背景图，优先级大于default_user_background_image

	UserBackgroundImages *[]MixUserBackgroundImage `json:"user_background_images,omitempty"`
}

func (o MixParam) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MixParam struct{}"
	}

	return strings.Join([]string{"MixParam", string(data)}, " ")
}

type MixParamEncodeTemplate struct {
	value string
}

type MixParamEncodeTemplateEnum struct {
	E_1920X1080_30_4620 MixParamEncodeTemplate
	E_1920X1080_15_3460 MixParamEncodeTemplate
	E_1280X720_30_3420  MixParamEncodeTemplate
	E_1280X720_15_2260  MixParamEncodeTemplate
}

func GetMixParamEncodeTemplateEnum() MixParamEncodeTemplateEnum {
	return MixParamEncodeTemplateEnum{
		E_1920X1080_30_4620: MixParamEncodeTemplate{
			value: "1920x1080_30_4620",
		},
		E_1920X1080_15_3460: MixParamEncodeTemplate{
			value: "1920x1080_15_3460",
		},
		E_1280X720_30_3420: MixParamEncodeTemplate{
			value: "1280x720_30_3420",
		},
		E_1280X720_15_2260: MixParamEncodeTemplate{
			value: "1280x720_15_2260",
		},
	}
}

func (c MixParamEncodeTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *MixParamEncodeTemplate) UnmarshalJSON(b []byte) error {
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

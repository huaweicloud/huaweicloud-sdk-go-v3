package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 录制视频混流参数。  用法如下：  如果某窗格需要展示某个固定用户的视频：  例如:{ \"id\": 1, \"user_id\": \"user001\", \"video_type\": \"CAMERASTREAM\" }  如果某窗格需要展示某个固定用户的共享屏幕：  例如:{ \"id\": 1, \"user_id\": \"user001\", \"video_type\": \"SCREENSTREAM\" }  如果某窗格需要展示共享桌面，不特定用户：  例如:{ \"id\": 1,  \"video_type\": \"SCREENSTREAM\"}
type MixLayoutPane struct {
	// 窗口id，从1开始编号

	Id *int32 `json:"id,omitempty"`
	// 加入房间的用户id

	UserId *string `json:"user_id,omitempty"`
	// 标识视频流的类型，可选摄像头流或者屏幕分享流。  - CAMERASTREAM：摄像头视频流 - SCREENSTREAM：屏幕分享视频流  默认为CAMERASTREAM。

	VideoType *MixLayoutPaneVideoType `json:"video_type,omitempty"`
	// 坐标x，归一化百分比，画布上该画面左上角的横坐标的相对值，范围是 [0.0,1.0]。从左到右布局，0.0在最左端，1.0在最右端，小数取值范围在float内，自定义布局场景下填写本字段。

	X *float32 `json:"x,omitempty"`
	// 坐标y，归一化百分比，画布上该画面左上角的纵坐标的相对值，范围是 [0.0,1.0]。从上到下布局，0.0在最上端，1.0在最下端，小数取值范围在float内，自定义布局场景下填写本字段。

	Y *float32 `json:"y,omitempty"`
	// 窗格宽，归一化百分比，小数取值范围在float内，自定义布局场景下填写本字段。

	Width *float32 `json:"width,omitempty"`
	// 窗格宽，归一化百分比，小数取值范围在float内，自定义布局场景下填写本字段。

	Height *float32 `json:"height,omitempty"`
	// 叠放顺序，0为最底层，1层在0层之上，以此类推，最大支持25层，自定义布局场景下填写本字段。

	Zorder *int32 `json:"zorder,omitempty"`
	// 裁剪模式，自定义布局场景下填写本字段，支持两种模式：   - KEEP_RATIO_PADDING ：保持比例留边。   - KEEP_RATIO_CROP ：保持比例裁剪。

	CropMode *MixLayoutPaneCropMode `json:"crop_mode,omitempty"`
}

func (o MixLayoutPane) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MixLayoutPane struct{}"
	}

	return strings.Join([]string{"MixLayoutPane", string(data)}, " ")
}

type MixLayoutPaneVideoType struct {
	value string
}

type MixLayoutPaneVideoTypeEnum struct {
	CAMERASTREAM MixLayoutPaneVideoType
	SCREENSTREAM MixLayoutPaneVideoType
}

func GetMixLayoutPaneVideoTypeEnum() MixLayoutPaneVideoTypeEnum {
	return MixLayoutPaneVideoTypeEnum{
		CAMERASTREAM: MixLayoutPaneVideoType{
			value: "CAMERASTREAM",
		},
		SCREENSTREAM: MixLayoutPaneVideoType{
			value: "SCREENSTREAM",
		},
	}
}

func (c MixLayoutPaneVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MixLayoutPaneVideoType) UnmarshalJSON(b []byte) error {
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

type MixLayoutPaneCropMode struct {
	value string
}

type MixLayoutPaneCropModeEnum struct {
	KEEP_RATIO_PADDING MixLayoutPaneCropMode
	KEEP_RATIO_CROP    MixLayoutPaneCropMode
}

func GetMixLayoutPaneCropModeEnum() MixLayoutPaneCropModeEnum {
	return MixLayoutPaneCropModeEnum{
		KEEP_RATIO_PADDING: MixLayoutPaneCropMode{
			value: "KEEP_RATIO_PADDING",
		},
		KEEP_RATIO_CROP: MixLayoutPaneCropMode{
			value: "KEEP_RATIO_CROP",
		},
	}
}

func (c MixLayoutPaneCropMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MixLayoutPaneCropMode) UnmarshalJSON(b []byte) error {
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

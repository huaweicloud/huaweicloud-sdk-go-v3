package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 切换会议显示策略请求体
type RestCustomMultiPictureBody struct {
	// 是否为手工设置多画面： 0： 系统自动多画面 1： 手工设置多画面

	ManualSet int32 `json:"manualSet"`

	PicLayoutInfo *PicLayoutInfo `json:"picLayoutInfo,omitempty"`
	// 画面类型

	ImageType *string `json:"imageType,omitempty"`
	// 子画面列表

	SubscriberInPics *[]RestSubscriberInPic `json:"subscriberInPics,omitempty"`
	// 表示轮询间隔，单位：秒。 当同一个子画面中包含有多个视频源时，此参数有效

	SwitchTime *int32 `json:"switchTime,omitempty"`
	// 多画面仅保存

	MultiPicSaveOnly *bool `json:"multiPicSaveOnly,omitempty"`
}

func (o RestCustomMultiPictureBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestCustomMultiPictureBody struct{}"
	}

	return strings.Join([]string{"RestCustomMultiPictureBody", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 当前多画面显示信息
type MultiPicDisplayDo struct {
	// 是否为手工设置多画面 0： 系统自动多画面 1： 手工设置多画面

	ManualSet *int32 `json:"manualSet,omitempty"`
	// 画面类型

	ImageType *string `json:"imageType,omitempty"`
	// 子画面列表

	SubscriberInPics *[]PicInfoNotify `json:"subscriberInPics,omitempty"`
	// 表示轮询间隔，单位：秒。当同一个子画面中包含有多个视频源时，此参数有效

	SwitchTime *string `json:"switchTime,omitempty"`

	PicLayoutInfo *PicLayoutInfo `json:"picLayoutInfo,omitempty"`
}

func (o MultiPicDisplayDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiPicDisplayDo struct{}"
	}

	return strings.Join([]string{"MultiPicDisplayDo", string(data)}, " ")
}

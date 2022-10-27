package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 多画面布局。
type RestPicLayout struct {

	// 多画面轮询时间，单位秒。
	SwitchTime *int32 `json:"switchTime,omitempty"`

	// 多画面画面数量。
	PicNum *int32 `json:"picNum,omitempty"`

	// 多画面布局名称。
	LayOutName *string `json:"layOutName,omitempty"`

	// 画面类型。
	ImageType *string `json:"imageType,omitempty"`

	// 布局UUID。
	Uuid *string `json:"uuid,omitempty"`

	// 子画面列表。
	SubscriberInPics *[]PicInfoNotify `json:"subscriberInPics,omitempty"`
}

func (o RestPicLayout) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestPicLayout struct{}"
	}

	return strings.Join([]string{"RestPicLayout", string(data)}, " ")
}

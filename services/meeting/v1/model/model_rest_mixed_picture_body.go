package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设置多画面请求。
type RestMixedPictureBody struct {

	// 是否为手工设置多画面。 - 0: 系统自动多画面 - 1: 手工设置多画面
	ManualSet int32 `json:"manualSet"`

	// 多画面数目。手工设置多画面时有效。 - Single: 单画面 - Two: 二画面 - Three: 三画面 - Four: 四画面 - Six: 六画面 - Nine: 九画面 - Sixteen: 十六画面
	ImageType *string `json:"imageType,omitempty"`

	// 子画面列表（手工设置多画面时必填）。
	SubscriberInPics *[]SubscriberInPic `json:"subscriberInPics,omitempty"`

	// 表示轮询间隔。单位：秒。 当同一个子画面中包含有多个与会者时，此参数有效。
	SwitchTime *int32 `json:"switchTime,omitempty"`
}

func (o RestMixedPictureBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestMixedPictureBody struct{}"
	}

	return strings.Join([]string{"RestMixedPictureBody", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestCustomMultiPictureBody 设置自定义多画面请求。
type RestCustomMultiPictureBody struct {

	// 是否为手工设置多画面。 * 0： 系统自动多画面 * 1： 手工设置多画面
	ManualSet int32 `json:"manualSet"`

	PicLayoutInfo *PicLayoutInfo `json:"picLayoutInfo,omitempty"`

	// 画面类型。手工设置多画面时有效。 - Single: 单画面 - Two: 二画面 - Three: 三画面 - Three-2: 三画面 - Three-3: 三画面 - Three-4: 三画面 - Four: 四画面 - Four-2: 四画面 - Four-3: 四画面 - Five: 五画面 - Five-2: 五画面 - Six: 六画面 - Six-2: 六画面 - Six-3: 六画面 - Six-4: 六画面 - Six-5: 六画面 - Seven: 七画面 - Seven-2: 七画面 - Seven-3: 七画面 - Seven-4: 七画面 - Eight: 八画面 - Eight-2: 八画面 - Eight-3: 八画面 - Eight-4: 八画面 - Nine: 九画面 - Ten: 十画面 - Ten-2: 十画面 - Ten-3: 十画面 - Ten-4: 十画面 - Ten-5: 十画面 - Ten-6: 十画面 - Thirteen: 十三画面 - Thirteen-2: 十三画面 - Thirteen-3: 十三画面 - Thirteen-4: 十三画面 - Thirteen-5: 十三画面 - Sixteen: 十六画面 - Seventeen: 十七画面 - Twenty-Five: 二十五画面
	ImageType *string `json:"imageType,omitempty"`

	// 子画面列表（手工设置多画面时必填）。
	SubscriberInPics *[]RestSubscriberInPic `json:"subscriberInPics,omitempty"`

	// 表示轮询间隔,取值范围：10-120，默认10。单位：秒。 当同一个子画面中包含有多个与会者时，此参数有效。 > 仅针对专业会议终端生效，对软终端不生效。
	SwitchTime *int32 `json:"switchTime,omitempty"`

	// 多画面是否仅保存。 * true： 仅保存 * false： 保存并应用 > * ”仅保存“效果：仅保存当前画面布局，不进行广播等操作。 > * ”保存并应用“效果：1、当会议状态为广播多画面、声控单画面、声控多画面、主持人观看多画面时，保存并应用后，改变画面布局，不改变状态。2、当会议状态为非广播多画面、声控单画面、声控多画面、主持人观看多画面时，如自动多画面、广播与会者、点名与会者时，保存并应用后，变为广播多画面。 > * 当处于广播多画面、声控多画面、声控单画面状态下，无法设置主持人观看多画面。
	MultiPicSaveOnly *bool `json:"multiPicSaveOnly,omitempty"`
}

func (o RestCustomMultiPictureBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestCustomMultiPictureBody struct{}"
	}

	return strings.Join([]string{"RestCustomMultiPictureBody", string(data)}, " ")
}

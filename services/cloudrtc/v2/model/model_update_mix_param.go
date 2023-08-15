package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMixParam 可修改的合流参数
type UpdateMixParam struct {

	// 视频布局模板编号，仅支持自定义模板之间的更新。
	LayoutTemplate *string `json:"layout_template,omitempty"`

	// 画布背景图地址，图片先上传obs。默认使用等比缩放裁剪，保证铺满。格式s3://bucket/object
	BackgroundImage *string `json:"background_image,omitempty"`

	// 默认用户背景图地址，图片先上传obs，格式s3://bucket/object。默认使用等比缩放裁剪，保证铺满。
	DefaultUserBackgroundImage *string `json:"default_user_background_image,omitempty"`

	// 共享屏幕的背景图地址，图片先上传obs，格式s3://bucket/object。  在一大多小的布局场景下，无论大窗是显示非指定用户（屏幕共享人的桌面）还是指定用户的共享桌面，都通过该字段指定背景图。
	ScreenBackgroundImage *string `json:"screen_background_image,omitempty"`

	// 最长空闲频道时间。  取值范围：[5，43200]，默认值为30。  单位：秒。  如果频道内无连麦方的状态持续超过该时间，录制程序会自动退出。退出后，再次调用start请求，会产生新的录制任务。  连麦方指：joiner或者publisher的用户。
	MaxIdleTime *int32 `json:"max_idle_time,omitempty"`

	// 需要混流的视频列表。若不需要视频混流，则可不传递该参数。
	LayoutPanes *[]MixLayoutPane `json:"layout_panes,omitempty"`

	// 指定用户背景图，优先级大于default_user_background_image
	UserBackgroundImages *[]MixUserBackgroundImage `json:"user_background_images,omitempty"`
}

func (o UpdateMixParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMixParam struct{}"
	}

	return strings.Join([]string{"UpdateMixParam", string(data)}, " ")
}

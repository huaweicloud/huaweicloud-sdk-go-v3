package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenRoomSettingVo 网络研讨会高级设置。
type OpenRoomSettingVo struct {

	// 网络研讨会介绍。
	RoomIntroduce *string `json:"roomIntroduce,omitempty"`

	// 企业Logo的文件id。
	CropLogoId *string `json:"cropLogoId,omitempty"`

	// 欢迎界面的文件id。
	CoverPictureId *string `json:"coverPictureId,omitempty"`

	// 显示观众人数的模式。默认值为real_time。 - none: 不显示 - real_time: 实时显示
	ShowAudienceMode *string `json:"showAudienceMode,omitempty"`

	// 智能倍增开关。默认值为Y。 - Y 开启智能倍增 - N 关闭智能倍增
	IsRedoubleOpen *string `json:"isRedoubleOpen,omitempty"`

	// 基础设置人数(网络研讨会没人时显示的人数)。默认值为0。取值范围为[0, 10000]。
	BaseAudienceCount *int32 `json:"baseAudienceCount,omitempty"`

	// 设置倍数(基础人数+真实人数*倍数)。默认值为1.0。 取值范围为[0, 10]，取1位小数。
	Multiple *float64 `json:"multiple,omitempty"`
}

func (o OpenRoomSettingVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenRoomSettingVo struct{}"
	}

	return strings.Join([]string{"OpenRoomSettingVo", string(data)}, " ")
}

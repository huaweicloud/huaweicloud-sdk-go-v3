package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRoomSettingResponse struct {

	// 直播介绍, 最大不超过500个字符
	RoomIntroduce *string `json:"roomIntroduce,omitempty"`

	// 企业Logo（文件id）,不超过32字符
	CropLogoId *string `json:"cropLogoId,omitempty"`

	// 封面内容(文件id)，不超过32字符
	CoverPictureId *string `json:"coverPictureId,omitempty"`

	// 显示观众人数的模式。默认值为real_time - none:不显示 - real_time:实时显示
	ShowAudienceMode *string `json:"showAudienceMode,omitempty"`

	// 智能倍增开关。默认值为Y - Y 开启智能倍增 - N 关闭智能倍增
	IsRedoubleOpen *string `json:"isRedoubleOpen,omitempty"`

	// 基础设置人数(直播间没人时显示的人数). 取值范围为[0, 10000]。默认值为0
	BaseAudienceCount *int32 `json:"baseAudienceCount,omitempty"`

	// 设置倍数(基础人数+真实人数*倍数). 取值范围为[0, 10]，取1位小数。默认值为1.0
	Multiple       *float64 `json:"multiple,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowRoomSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRoomSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowRoomSettingResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSmartLiveRoomResponse Response Object
type UpdateSmartLiveRoomResponse struct {

	// 直播间名称
	RoomName string `json:"room_name"`

	// 直播间描述。
	RoomDescription *string `json:"room_description,omitempty"`

	// 直播间类型。 * NORMAL: 普通直播间，直播间一直存在，可以反复开播 * TEMP: 临时直播间,直播任务结束后自动清理直播间。
	RoomType *UpdateSmartLiveRoomResponseRoomType `json:"room_type,omitempty"`

	// 默认直播剧本列表。
	SceneScripts *[]LiveVideoScriptInfo `json:"scene_scripts,omitempty"`

	// 互动规则列表
	InteractionRules *[]InteractionRuleInfo `json:"interaction_rules,omitempty"`

	PlayPolicy *PlayPolicy `json:"play_policy,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 视频推流第三方直播平台地址。
	OutputUrls *[]string `json:"output_urls,omitempty"`

	// 直播间ID
	RoomId *string `json:"room_id,omitempty"`

	// 直播间创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 直播间更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	// 直播间封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSmartLiveRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSmartLiveRoomResponse struct{}"
	}

	return strings.Join([]string{"UpdateSmartLiveRoomResponse", string(data)}, " ")
}

type UpdateSmartLiveRoomResponseRoomType struct {
	value string
}

type UpdateSmartLiveRoomResponseRoomTypeEnum struct {
	NORMAL UpdateSmartLiveRoomResponseRoomType
	TEMP   UpdateSmartLiveRoomResponseRoomType
}

func GetUpdateSmartLiveRoomResponseRoomTypeEnum() UpdateSmartLiveRoomResponseRoomTypeEnum {
	return UpdateSmartLiveRoomResponseRoomTypeEnum{
		NORMAL: UpdateSmartLiveRoomResponseRoomType{
			value: "NORMAL",
		},
		TEMP: UpdateSmartLiveRoomResponseRoomType{
			value: "TEMP",
		},
	}
}

func (c UpdateSmartLiveRoomResponseRoomType) Value() string {
	return c.value
}

func (c UpdateSmartLiveRoomResponseRoomType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSmartLiveRoomResponseRoomType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

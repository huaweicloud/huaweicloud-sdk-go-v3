package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSmartLiveRoomReq 创建直播间配置。
type CreateSmartLiveRoomReq struct {

	// 直播间名称
	RoomName string `json:"room_name"`

	// 直播间描述。
	RoomDescription *string `json:"room_description,omitempty"`

	// 直播间类型。 * NORMAL: 普通直播间，直播间一直存在，可以反复开播 * TEMP: 临时直播间,直播任务结束后自动清理直播间。
	RoomType *CreateSmartLiveRoomReqRoomType `json:"room_type,omitempty"`

	// 默认直播剧本列表。
	SceneScripts *[]LiveVideoScriptInfo `json:"scene_scripts,omitempty"`

	// 互动规则列表
	InteractionRules *[]InteractionRuleInfo `json:"interaction_rules,omitempty"`

	PlayPolicy *PlayPolicy `json:"play_policy,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 视频推流第三方直播平台地址。
	OutputUrls *[]string `json:"output_urls,omitempty"`
}

func (o CreateSmartLiveRoomReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSmartLiveRoomReq struct{}"
	}

	return strings.Join([]string{"CreateSmartLiveRoomReq", string(data)}, " ")
}

type CreateSmartLiveRoomReqRoomType struct {
	value string
}

type CreateSmartLiveRoomReqRoomTypeEnum struct {
	NORMAL CreateSmartLiveRoomReqRoomType
	TEMP   CreateSmartLiveRoomReqRoomType
}

func GetCreateSmartLiveRoomReqRoomTypeEnum() CreateSmartLiveRoomReqRoomTypeEnum {
	return CreateSmartLiveRoomReqRoomTypeEnum{
		NORMAL: CreateSmartLiveRoomReqRoomType{
			value: "NORMAL",
		},
		TEMP: CreateSmartLiveRoomReqRoomType{
			value: "TEMP",
		},
	}
}

func (c CreateSmartLiveRoomReqRoomType) Value() string {
	return c.value
}

func (c CreateSmartLiveRoomReqRoomType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSmartLiveRoomReqRoomType) UnmarshalJSON(b []byte) error {
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

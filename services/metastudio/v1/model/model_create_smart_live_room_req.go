package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSmartLiveRoomReq 创建直播间配置。
type CreateSmartLiveRoomReq struct {

	// **参数解释**： 直播间名称。 **约束限制**： 不涉及。 **取值范围**： 字符长度1-256位。 **默认取值**： 不涉及。
	RoomName string `json:"room_name"`

	// **参数解释**： 直播间描述。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-1024位。 **默认取值**： 不涉及。
	RoomDescription *string `json:"room_description,omitempty"`

	// **参数解释**： 直播间类型。 **约束限制**： 不涉及。 **取值范围**： * NORMAL：普通直播间，直播间一直存在，可以反复开播 * TEMP：临时直播间，直播任务结束后自动清理直播间。 * TEMPLATE：直播间模板。
	RoomType *CreateSmartLiveRoomReqRoomType `json:"room_type,omitempty"`

	// 默认直播剧本列表。
	SceneScripts *[]LiveVideoScriptInfo `json:"scene_scripts,omitempty"`

	InteractionConfig *LiveRoomInteractionConfig `json:"interaction_config,omitempty"`

	// 互动规则列表
	InteractionRules *[]LiveRoomInteractionRuleInfo `json:"interaction_rules,omitempty"`

	PlayPolicy *PlayPolicy `json:"play_policy,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// **参数解释**： RTMP视频推流第三方直播平台地址。 > 直播过程中刷新地址，需要调用COMMAND命令REFRESH_OUTPUT_URL。  **约束限制**： 不涉及 **取值范围**： 当前仅支持一条RTMP出流地址。 **默认取值**： 不涉及。
	OutputUrls *[]string `json:"output_urls,omitempty"`

	// **参数解释**： RTMP视频推流第三方直播平台流密钥，与推流地址对应。 > 直播过程中刷新地址，需要调用COMMAND命令REFRESH_OUTPUT_URL。  **约束限制**： 不涉及 **取值范围**： 当前仅支持一条RTMP出流地址。 **默认取值**： 不涉及。
	StreamKeys *[]string `json:"stream_keys,omitempty"`

	// **参数解释**： 主播轮换时备选主播数字人资产ID（仅形象资产，不包含声音）。  **约束限制**： 不涉及 **取值范围**： 当前最大支持5个备选主播。 数字人资产ID，字符长度0-64位。 **默认取值**： 不涉及
	BackupModelAssetIds *[]string `json:"backup_model_asset_ids,omitempty"`

	LiveEventCallbackConfig *LiveEventCallBackConfig `json:"live_event_callback_config,omitempty"`

	RtcCallbackConfig *RtcLiveEventCallBackConfig `json:"rtc_callback_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	SharedConfig *SharedConfig `json:"shared_config,omitempty"`

	// **参数解释**： 横竖屏类型。 **约束限制**： 用户无需填写，通过video_config中分辨率判断 **取值范围**： * LANDSCAPE：横屏。 * VERTICAL： 竖屏。
	ViewMode *CreateSmartLiveRoomReqViewMode `json:"view_mode,omitempty"`

	CoStreamerConfig *CoStreamerConfig `json:"co_streamer_config,omitempty"`

	// **参数解释**： 匹配值私有数据，用户填写，原样带回。 **约束限制**： 不涉及 **取值范围**： 字符长度0-8192 **默认取值**： 不涉及。
	PrivData *string `json:"priv_data,omitempty"`
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
	NORMAL   CreateSmartLiveRoomReqRoomType
	TEMP     CreateSmartLiveRoomReqRoomType
	TEMPLATE CreateSmartLiveRoomReqRoomType
}

func GetCreateSmartLiveRoomReqRoomTypeEnum() CreateSmartLiveRoomReqRoomTypeEnum {
	return CreateSmartLiveRoomReqRoomTypeEnum{
		NORMAL: CreateSmartLiveRoomReqRoomType{
			value: "NORMAL",
		},
		TEMP: CreateSmartLiveRoomReqRoomType{
			value: "TEMP",
		},
		TEMPLATE: CreateSmartLiveRoomReqRoomType{
			value: "TEMPLATE",
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

type CreateSmartLiveRoomReqViewMode struct {
	value string
}

type CreateSmartLiveRoomReqViewModeEnum struct {
	LANDSCAPE CreateSmartLiveRoomReqViewMode
	VERTICAL  CreateSmartLiveRoomReqViewMode
}

func GetCreateSmartLiveRoomReqViewModeEnum() CreateSmartLiveRoomReqViewModeEnum {
	return CreateSmartLiveRoomReqViewModeEnum{
		LANDSCAPE: CreateSmartLiveRoomReqViewMode{
			value: "LANDSCAPE",
		},
		VERTICAL: CreateSmartLiveRoomReqViewMode{
			value: "VERTICAL",
		},
	}
}

func (c CreateSmartLiveRoomReqViewMode) Value() string {
	return c.value
}

func (c CreateSmartLiveRoomReqViewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSmartLiveRoomReqViewMode) UnmarshalJSON(b []byte) error {
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

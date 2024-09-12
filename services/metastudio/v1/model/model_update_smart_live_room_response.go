package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSmartLiveRoomResponse Response Object
type UpdateSmartLiveRoomResponse struct {

	// **参数解释**： 直播间名称。 **约束限制**： 不涉及。 **取值范围**： 字符长度1-256位。 **默认取值**： 不涉及。
	RoomName string `json:"room_name"`

	// **参数解释**： 直播间描述。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-1024位。 **默认取值**： 不涉及。
	RoomDescription *string `json:"room_description,omitempty"`

	// **参数解释**： 直播间类型。 **约束限制**： 不涉及。 **取值范围**： * NORMAL：普通直播间，直播间一直存在，可以反复开播 * TEMP：临时直播间，直播任务结束后自动清理直播间。 * TEMPLATE：直播间模板。
	RoomType *UpdateSmartLiveRoomResponseRoomType `json:"room_type,omitempty"`

	// 默认直播剧本列表。
	SceneScripts *[]LiveVideoScriptInfo `json:"scene_scripts,omitempty"`

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
	ViewMode *UpdateSmartLiveRoomResponseViewMode `json:"view_mode,omitempty"`

	CoStreamerConfig *CoStreamerConfig `json:"co_streamer_config,omitempty"`

	// **参数解释**： 匹配值私有数据，用户填写，原样带回。 **约束限制**： 不涉及 **取值范围**： 字符长度0-8192 **默认取值**： 不涉及。
	PrivData *string `json:"priv_data,omitempty"`

	// 直播间ID
	RoomId *string `json:"room_id,omitempty"`

	// 直播间创建时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 直播间更新时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`

	// 直播间封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`

	// 直播间封面图新URL
	Thumbnail *string `json:"thumbnail,omitempty"`

	// 直播间配置状态。 - ENABLE: 直播间正常可用。 - DISABLE： 直播间不可用。不可用原因在error_info中说明。 - BLOCKED：直播间被冻结。冻结原因在error_info中说明。
	RoomState *UpdateSmartLiveRoomResponseRoomState `json:"room_state,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

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
	NORMAL   UpdateSmartLiveRoomResponseRoomType
	TEMP     UpdateSmartLiveRoomResponseRoomType
	TEMPLATE UpdateSmartLiveRoomResponseRoomType
}

func GetUpdateSmartLiveRoomResponseRoomTypeEnum() UpdateSmartLiveRoomResponseRoomTypeEnum {
	return UpdateSmartLiveRoomResponseRoomTypeEnum{
		NORMAL: UpdateSmartLiveRoomResponseRoomType{
			value: "NORMAL",
		},
		TEMP: UpdateSmartLiveRoomResponseRoomType{
			value: "TEMP",
		},
		TEMPLATE: UpdateSmartLiveRoomResponseRoomType{
			value: "TEMPLATE",
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

type UpdateSmartLiveRoomResponseViewMode struct {
	value string
}

type UpdateSmartLiveRoomResponseViewModeEnum struct {
	LANDSCAPE UpdateSmartLiveRoomResponseViewMode
	VERTICAL  UpdateSmartLiveRoomResponseViewMode
}

func GetUpdateSmartLiveRoomResponseViewModeEnum() UpdateSmartLiveRoomResponseViewModeEnum {
	return UpdateSmartLiveRoomResponseViewModeEnum{
		LANDSCAPE: UpdateSmartLiveRoomResponseViewMode{
			value: "LANDSCAPE",
		},
		VERTICAL: UpdateSmartLiveRoomResponseViewMode{
			value: "VERTICAL",
		},
	}
}

func (c UpdateSmartLiveRoomResponseViewMode) Value() string {
	return c.value
}

func (c UpdateSmartLiveRoomResponseViewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSmartLiveRoomResponseViewMode) UnmarshalJSON(b []byte) error {
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

type UpdateSmartLiveRoomResponseRoomState struct {
	value string
}

type UpdateSmartLiveRoomResponseRoomStateEnum struct {
	ENABLE  UpdateSmartLiveRoomResponseRoomState
	DISABLE UpdateSmartLiveRoomResponseRoomState
	BLOCKED UpdateSmartLiveRoomResponseRoomState
}

func GetUpdateSmartLiveRoomResponseRoomStateEnum() UpdateSmartLiveRoomResponseRoomStateEnum {
	return UpdateSmartLiveRoomResponseRoomStateEnum{
		ENABLE: UpdateSmartLiveRoomResponseRoomState{
			value: "ENABLE",
		},
		DISABLE: UpdateSmartLiveRoomResponseRoomState{
			value: "DISABLE",
		},
		BLOCKED: UpdateSmartLiveRoomResponseRoomState{
			value: "BLOCKED",
		},
	}
}

func (c UpdateSmartLiveRoomResponseRoomState) Value() string {
	return c.value
}

func (c UpdateSmartLiveRoomResponseRoomState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSmartLiveRoomResponseRoomState) UnmarshalJSON(b []byte) error {
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

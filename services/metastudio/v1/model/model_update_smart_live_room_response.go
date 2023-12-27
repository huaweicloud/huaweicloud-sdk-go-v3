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

	// 直播间类型。 * NORMAL: 普通直播间，直播间一直存在，可以反复开播 * TEMP: 临时直播间,直播任务结束后自动清理直播间。 * TEMPLATE: 直播间模板。
	RoomType *UpdateSmartLiveRoomResponseRoomType `json:"room_type,omitempty"`

	// 默认直播剧本列表。
	SceneScripts *[]LiveVideoScriptInfo `json:"scene_scripts,omitempty"`

	// 互动规则列表
	InteractionRules *[]InteractionRuleInfo `json:"interaction_rules,omitempty"`

	PlayPolicy *PlayPolicy `json:"play_policy,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// RTMP视频推流第三方直播平台地址。
	OutputUrls *[]string `json:"output_urls,omitempty"`

	// RTMP视频推流第三方直播平台流秘钥，与推流地址对应。
	StreamKeys *[]string `json:"stream_keys,omitempty"`

	// 主播轮换时备选主播数字人资产ID（仅形象资产，不包含音色）。
	BackupModelAssetIds *[]string `json:"backup_model_asset_ids,omitempty"`

	LiveEventCallbackConfig *LiveEventCallBackConfig `json:"live_event_callback_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	SharedConfig *SharedConfig `json:"shared_config,omitempty"`

	// 直播间ID
	RoomId *string `json:"room_id,omitempty"`

	// 直播间创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 直播间更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	// 直播间封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`

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

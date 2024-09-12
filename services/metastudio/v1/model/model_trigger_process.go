package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TriggerProcess 触发器处理
type TriggerProcess struct {

	// **参数解释**： 处理抑制时长。单位秒。 - -1：表示整场直播仅触发一次。 - 0：表示无抑制，每次都触发。 - 其他值n：表示n秒内仅触发一次。  **约束限制**： 不涉及 **默认取值**： 不涉及
	TimeWindow *int32 `json:"time_window,omitempty"`

	// **参数解释**： 回复类型。 **约束限制**： 不涉及 **取值范围**： * SYSTEM_REPLY：系统自动回复预先设置的话术。 * CALLBACK：回调给其他服务，携带设置的话术。 * SHOW_LAYER：仅显示叠加图层，不影响话术。 * INTELLIGENT_REPLY：使用配置的大模型生成回复话术。  **默认取值**： 不涉及
	ReplyMode *TriggerProcessReplyMode `json:"reply_mode,omitempty"`

	LayerConfig *SmartLayerConfig `json:"layer_config,omitempty"`

	ExtraLayerConfig *SmartLayerConfig `json:"extra_layer_config,omitempty"`

	// **参数解释**： 回复话术集。 **约束限制**： 不涉及 **取值范围**： 最大支持5条预置话术。 单条话术字符长度0-1024位。 **默认取值**： 不涉及
	ReplyTexts *[]string `json:"reply_texts,omitempty"`

	// **参数解释**： 回复音频集。填写audio_url。 **约束限制**： 不涉及 **取值范围**： 最大支持5条预置音频。 **默认取值**： 不涉及
	ReplyAudios *[]ReplyAudioInfo `json:"reply_audios,omitempty"`

	// **参数解释**： 回复话术选择次序。 **约束限制**： 不涉及 **取值范围**： * RANDOM：随机 * ORDER：顺序循环  **默认取值**： 不涉及
	ReplyOrder *TriggerProcessReplyOrder `json:"reply_order,omitempty"`

	// **参数解释**： 回复角色。 **约束限制**： 不涉及 **取值范围**： * STREAMER：主播 * CO_STREAMER：助播，仅声音。
	ReplyRole *TriggerProcessReplyRole `json:"reply_role,omitempty"`

	// **参数解释**： 机器人ID。 **约束限制**： reply_mode为INTELLIGENT_REPLY时必填，智能交互配置的大模型机器人ID。 获取方法请参考[创建应用](CreateRobot.xml)。 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及
	RobotId *string `json:"robot_id,omitempty"`
}

func (o TriggerProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerProcess struct{}"
	}

	return strings.Join([]string{"TriggerProcess", string(data)}, " ")
}

type TriggerProcessReplyMode struct {
	value string
}

type TriggerProcessReplyModeEnum struct {
	SYSTEM_REPLY      TriggerProcessReplyMode
	CALLBACK          TriggerProcessReplyMode
	SHOW_LAYER        TriggerProcessReplyMode
	INTELLIGENT_REPLY TriggerProcessReplyMode
}

func GetTriggerProcessReplyModeEnum() TriggerProcessReplyModeEnum {
	return TriggerProcessReplyModeEnum{
		SYSTEM_REPLY: TriggerProcessReplyMode{
			value: "SYSTEM_REPLY",
		},
		CALLBACK: TriggerProcessReplyMode{
			value: "CALLBACK",
		},
		SHOW_LAYER: TriggerProcessReplyMode{
			value: "SHOW_LAYER",
		},
		INTELLIGENT_REPLY: TriggerProcessReplyMode{
			value: "INTELLIGENT_REPLY",
		},
	}
}

func (c TriggerProcessReplyMode) Value() string {
	return c.value
}

func (c TriggerProcessReplyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerProcessReplyMode) UnmarshalJSON(b []byte) error {
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

type TriggerProcessReplyOrder struct {
	value string
}

type TriggerProcessReplyOrderEnum struct {
	RANDOM TriggerProcessReplyOrder
	ORDER  TriggerProcessReplyOrder
}

func GetTriggerProcessReplyOrderEnum() TriggerProcessReplyOrderEnum {
	return TriggerProcessReplyOrderEnum{
		RANDOM: TriggerProcessReplyOrder{
			value: "RANDOM",
		},
		ORDER: TriggerProcessReplyOrder{
			value: "ORDER",
		},
	}
}

func (c TriggerProcessReplyOrder) Value() string {
	return c.value
}

func (c TriggerProcessReplyOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerProcessReplyOrder) UnmarshalJSON(b []byte) error {
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

type TriggerProcessReplyRole struct {
	value string
}

type TriggerProcessReplyRoleEnum struct {
	STREAMER    TriggerProcessReplyRole
	CO_STREAMER TriggerProcessReplyRole
}

func GetTriggerProcessReplyRoleEnum() TriggerProcessReplyRoleEnum {
	return TriggerProcessReplyRoleEnum{
		STREAMER: TriggerProcessReplyRole{
			value: "STREAMER",
		},
		CO_STREAMER: TriggerProcessReplyRole{
			value: "CO_STREAMER",
		},
	}
}

func (c TriggerProcessReplyRole) Value() string {
	return c.value
}

func (c TriggerProcessReplyRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerProcessReplyRole) UnmarshalJSON(b []byte) error {
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

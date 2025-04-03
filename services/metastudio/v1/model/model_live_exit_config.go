package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LiveExitConfig 直播任务退出配置
type LiveExitConfig struct {

	// **参数解释**： 最大直播时长。单位小时。 此数值配置为n，则标识起播后n小时后触发停止直播逻辑。 当前数值最大为168（一周），特殊的，0表示不限时。 **约束限制**： 停止直播逻辑配置为立即停止则直播停止误差在5分钟之内。其他逻辑则加上处理时长。 **默认取值**： 不设置则表示不限时。
	MaxLiveDuration *int32 `json:"max_live_duration,omitempty"`

	// **参数解释**： 自动停止直播模式。 * FORCE_EXIT：立即强制停止，不做其他逻辑处理。 * SEGMENT_END:等待段落结束停止。 * SCENE_END：等待场景结束停止。 * SCRIPT_END：等待剧本结束停止。 **约束限制**： 不涉及 **默认取值**： 不设置则表示FORCE_EXIT。
	AutoStopMode *LiveExitConfigAutoStopMode `json:"auto_stop_mode,omitempty"`

	// **参数解释**： 最大异常等待时长。单位分钟。  此数值配置为n，则标识检测到异常n分钟后触发立即停止直播逻辑。 当前数值最大为60（1小时），特殊的，0表示不限时。 **约束限制**： 不涉及 **默认取值**： 不设置则为系统默认值，当前为3分钟，默认值可能会根据服务运行状态进行少许调整。
	MaxExceptionWaitingDuration *int32 `json:"max_exception_waiting_duration,omitempty"`
}

func (o LiveExitConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveExitConfig struct{}"
	}

	return strings.Join([]string{"LiveExitConfig", string(data)}, " ")
}

type LiveExitConfigAutoStopMode struct {
	value string
}

type LiveExitConfigAutoStopModeEnum struct {
	FORCE_EXIT  LiveExitConfigAutoStopMode
	SEGMENT_END LiveExitConfigAutoStopMode
	SCENE_END   LiveExitConfigAutoStopMode
	SCRIPT_END  LiveExitConfigAutoStopMode
}

func GetLiveExitConfigAutoStopModeEnum() LiveExitConfigAutoStopModeEnum {
	return LiveExitConfigAutoStopModeEnum{
		FORCE_EXIT: LiveExitConfigAutoStopMode{
			value: "FORCE_EXIT",
		},
		SEGMENT_END: LiveExitConfigAutoStopMode{
			value: "SEGMENT_END",
		},
		SCENE_END: LiveExitConfigAutoStopMode{
			value: "SCENE_END",
		},
		SCRIPT_END: LiveExitConfigAutoStopMode{
			value: "SCRIPT_END",
		},
	}
}

func (c LiveExitConfigAutoStopMode) Value() string {
	return c.value
}

func (c LiveExitConfigAutoStopMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveExitConfigAutoStopMode) UnmarshalJSON(b []byte) error {
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

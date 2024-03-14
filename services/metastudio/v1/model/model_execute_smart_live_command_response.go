package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteSmartLiveCommandResponse Response Object
type ExecuteSmartLiveCommandResponse struct {

	// 控制命令ID
	CommandId *string `json:"command_id,omitempty"`

	// 命令名称。 - INSERT_PLAY_SCRIPT: 插入表演脚本。用于互动回复。数字人不变，背景不变。params结构定义：ShootScript - REWRITE_PLAY_SCRIPT: 动态编辑未播放剧本。params结构定义：scene_scripts - INSERT_PLAY_AUDIO: 插入驱动音频。用于音频直接驱动。数字人不变，背景不变。params结构定义：PlayAudioInfo - GET_CURRENT_PLAYING_SCRIPTS: 查询本轮剧本列表。响应为LivePlayingScriptList结构 - REWRITE_INTERACTION_RULES: 动态修改互动规则。params结构定义：interaction_rules - GET_LIVE_JOB_CONFIG_INFO: 获取任务中的房间信息。params结构定义：SmartLiveRoomInfo
	Command *ExecuteSmartLiveCommandResponseCommand `json:"command,omitempty"`

	// 命令执行结果
	Result *string `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteSmartLiveCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSmartLiveCommandResponse struct{}"
	}

	return strings.Join([]string{"ExecuteSmartLiveCommandResponse", string(data)}, " ")
}

type ExecuteSmartLiveCommandResponseCommand struct {
	value string
}

type ExecuteSmartLiveCommandResponseCommandEnum struct {
	INSERT_PLAY_SCRIPT          ExecuteSmartLiveCommandResponseCommand
	REWRITE_PLAY_SCRIPT         ExecuteSmartLiveCommandResponseCommand
	INSERT_PLAY_AUDIO           ExecuteSmartLiveCommandResponseCommand
	GET_CURRENT_PLAYING_SCRIPTS ExecuteSmartLiveCommandResponseCommand
	REWRITE_INTERACTION_RULES   ExecuteSmartLiveCommandResponseCommand
	GET_LIVE_JOB_CONFIG_INFO    ExecuteSmartLiveCommandResponseCommand
}

func GetExecuteSmartLiveCommandResponseCommandEnum() ExecuteSmartLiveCommandResponseCommandEnum {
	return ExecuteSmartLiveCommandResponseCommandEnum{
		INSERT_PLAY_SCRIPT: ExecuteSmartLiveCommandResponseCommand{
			value: "INSERT_PLAY_SCRIPT",
		},
		REWRITE_PLAY_SCRIPT: ExecuteSmartLiveCommandResponseCommand{
			value: "REWRITE_PLAY_SCRIPT",
		},
		INSERT_PLAY_AUDIO: ExecuteSmartLiveCommandResponseCommand{
			value: "INSERT_PLAY_AUDIO",
		},
		GET_CURRENT_PLAYING_SCRIPTS: ExecuteSmartLiveCommandResponseCommand{
			value: "GET_CURRENT_PLAYING_SCRIPTS",
		},
		REWRITE_INTERACTION_RULES: ExecuteSmartLiveCommandResponseCommand{
			value: "REWRITE_INTERACTION_RULES",
		},
		GET_LIVE_JOB_CONFIG_INFO: ExecuteSmartLiveCommandResponseCommand{
			value: "GET_LIVE_JOB_CONFIG_INFO",
		},
	}
}

func (c ExecuteSmartLiveCommandResponseCommand) Value() string {
	return c.value
}

func (c ExecuteSmartLiveCommandResponseCommand) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteSmartLiveCommandResponseCommand) UnmarshalJSON(b []byte) error {
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

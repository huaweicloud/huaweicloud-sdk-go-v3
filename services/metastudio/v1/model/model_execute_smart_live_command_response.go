package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteSmartLiveCommandResponse Response Object
type ExecuteSmartLiveCommandResponse struct {

	// 命令名称。 - INSERT_PLAY_SCRIPT: 插入表演脚本。用于互动回复。数字人不变，背景不变。params结构定义：ShootScript - REWRITE_PLAY_SCRIPT: 动态编辑未播放剧本。params结构定义：scence_scripts - INSERT_PLAY_ADUIO: 插入驱动音频。用于音频直接驱动。数字人不变，背景不变。params结构定义：PlayAudioInfo - GET_CURRENT_PLAYING_SCRIPTS: 查询本轮剧本列表。响应为LivePlayingScriptList结构
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

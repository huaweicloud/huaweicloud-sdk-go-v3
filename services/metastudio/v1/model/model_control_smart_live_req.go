package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ControlSmartLiveReq 控制命令。
type ControlSmartLiveReq struct {

	// **参数解释**： 命令名称。 **约束限制**： 不限制 **取值范围**： * INSERT_PLAY_SCRIPT：插入表演脚本。用于互动回复。数字人不变，背景不变。params结构定义：[PlayTextInfo](metastudio_02_0014.xml#section0)。 * INSERT_PLAY_AUDIO：插入驱动音频。用于音频直接驱动。数字人不变，背景不变。params结构定义：[PlayAudioInfo](metastudio_02_0014.xml#section1)。 * REWRITE_PLAY_SCRIPT：动态编辑未播放剧本。params结构定义：[scene_scripts](CreateSmartLiveRoom.xml)。 * REWRITE_INTERACTION_RULES：动态修改互动规则。params结构定义：[interaction_rules](CreateSmartLiveRoom.xml)。 * GET_CURRENT_PLAYING_SCRIPTS：查询本轮剧本列表。响应为：[LivePlayingScriptList](metastudio_02_0014.xml#section2)结构。 * SHOW_LAYER：显示导播素材，用于直播导播。params结构定义：LiveGuideRuleInfo。 * REFRESH_OUTPUT_URL：更新当前任务的rtmp推流信息。params结构定义： RefreshOutputUrlConfig。 * GET_LIVE_JOB_CONFIG_INFO：获取任务中的房间信息。params结构定义：与[直播间详情响应体](ShowSmartLiveRoom.xml)一致。 * CLEAN_UP_INSERT_COMMAND：清理未播放的插入命令。params结构定义：[CleanUpInsertCommand](metastudio_02_0014.xml#section3) **默认取值**： 不涉及
	Command ControlSmartLiveReqCommand `json:"command"`

	// **参数解释**： 命令参数。 **约束限制**： 不限制 **取值范围**： 参考COMMNAD说明。 **默认取值**： 不涉及
	Params *interface{} `json:"params,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`
}

func (o ControlSmartLiveReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ControlSmartLiveReq struct{}"
	}

	return strings.Join([]string{"ControlSmartLiveReq", string(data)}, " ")
}

type ControlSmartLiveReqCommand struct {
	value string
}

type ControlSmartLiveReqCommandEnum struct {
	INSERT_PLAY_SCRIPT          ControlSmartLiveReqCommand
	REWRITE_PLAY_SCRIPT         ControlSmartLiveReqCommand
	INSERT_PLAY_AUDIO           ControlSmartLiveReqCommand
	REWRITE_INTERACTION_RULES   ControlSmartLiveReqCommand
	GET_CURRENT_PLAYING_SCRIPTS ControlSmartLiveReqCommand
	REFRESH_OUTPUT_URL          ControlSmartLiveReqCommand
	GET_LIVE_JOB_CONFIG_INFO    ControlSmartLiveReqCommand
	CLEAN_UP_INSERT_COMMAND     ControlSmartLiveReqCommand
}

func GetControlSmartLiveReqCommandEnum() ControlSmartLiveReqCommandEnum {
	return ControlSmartLiveReqCommandEnum{
		INSERT_PLAY_SCRIPT: ControlSmartLiveReqCommand{
			value: "INSERT_PLAY_SCRIPT",
		},
		REWRITE_PLAY_SCRIPT: ControlSmartLiveReqCommand{
			value: "REWRITE_PLAY_SCRIPT",
		},
		INSERT_PLAY_AUDIO: ControlSmartLiveReqCommand{
			value: "INSERT_PLAY_AUDIO",
		},
		REWRITE_INTERACTION_RULES: ControlSmartLiveReqCommand{
			value: "REWRITE_INTERACTION_RULES",
		},
		GET_CURRENT_PLAYING_SCRIPTS: ControlSmartLiveReqCommand{
			value: "GET_CURRENT_PLAYING_SCRIPTS",
		},
		REFRESH_OUTPUT_URL: ControlSmartLiveReqCommand{
			value: "REFRESH_OUTPUT_URL",
		},
		GET_LIVE_JOB_CONFIG_INFO: ControlSmartLiveReqCommand{
			value: "GET_LIVE_JOB_CONFIG_INFO",
		},
		CLEAN_UP_INSERT_COMMAND: ControlSmartLiveReqCommand{
			value: "CLEAN_UP_INSERT_COMMAND",
		},
	}
}

func (c ControlSmartLiveReqCommand) Value() string {
	return c.value
}

func (c ControlSmartLiveReqCommand) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ControlSmartLiveReqCommand) UnmarshalJSON(b []byte) error {
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

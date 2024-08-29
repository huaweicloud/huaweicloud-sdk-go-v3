package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRobotReq 修改应用请求。
type UpdateRobotReq struct {

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 对接第三方应用厂商类型。 > 0：科大讯飞AIUI；1：华为云CBS；2：科大讯飞星火交互认知大模型；5：第三方驱动；6：第三方语言模型；8：奇妙问
	AppType *int32 `json:"app_type,omitempty"`

	// 智能交互对话房间ID。
	RoomId *string `json:"room_id,omitempty"`

	RobotType *RobotTypeEnum `json:"robot_type,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	// 语音识别后端点静音时长默认500ms
	TailSilenceTime *int32 `json:"tail_silence_time,omitempty"`

	// 提问文本审核开关
	EnableQuestionAudit *bool `json:"enable_question_audit,omitempty"`

	HuaweiEiCbs *HuaweiEiCbs `json:"huawei_ei_cbs,omitempty"`

	IflytekAiuiConfig *IflytekAiuiConfig `json:"iflytek_aiui_config,omitempty"`

	IflytekSpark *IflytekSpark `json:"iflytek_spark,omitempty"`

	ThirdPartyModelConfig *ThirdPartyModelConfig `json:"third_party_model_config,omitempty"`

	MobvoiConfig *MobvoiConfig `json:"mobvoi_config,omitempty"`
}

func (o UpdateRobotReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRobotReq struct{}"
	}

	return strings.Join([]string{"UpdateRobotReq", string(data)}, " ")
}

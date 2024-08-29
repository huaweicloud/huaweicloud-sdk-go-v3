package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RobotInfo 应用基本信息。
type RobotInfo struct {

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 智能交互对话房间ID。
	RoomId *string `json:"room_id,omitempty"`

	// 第三方应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 对接第三方应用厂商类型。 > 0：科大讯飞AIUI；1：华为云CBS；2：科大讯飞星火交互认知大模型；5：第三方驱动；6：第三方语言模型；8：奇妙问
	AppType *int32 `json:"app_type,omitempty"`

	// 应用的AccessKey或帐号。
	AppKey *string `json:"app_key,omitempty"`

	RobotType *RobotTypeEnum `json:"robot_type,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	// CBS所在区域
	Region *int32 `json:"region,omitempty"`

	// CBS所在区域的projectId
	CbsProjectId *string `json:"cbs_project_id,omitempty"`

	// 第三方语言模型地址。
	LlmUrl *string `json:"llm_url,omitempty"`

	// 是否采用流式响应。
	IsStream *bool `json:"is_stream,omitempty"`

	// 支持的多轮对话数量，取值大于1时，请求第三方语言模型时将携带历史对话信息。
	ChatRounds *int32 `json:"chat_rounds,omitempty"`

	// 是否为正式环境
	IsIflyProduction *bool `json:"is_ifly_production,omitempty"`

	// 语音识别后端点静音时长默认500ms
	TailSilenceTime *int32 `json:"tail_silence_time,omitempty"`

	// 奇妙问角色ID。
	RoleId *string `json:"role_id,omitempty"`

	// SIS所在区域
	SisRegion *int32 `json:"sis_region,omitempty"`

	// SIS所在区域的projectId
	SisProjectId *string `json:"sis_project_id,omitempty"`

	// 是否开启热词
	EnableHotWords *bool `json:"enable_hot_words,omitempty"`

	// 是否开启提问文本审核开关
	EnableQuestionAudit *bool `json:"enable_question_audit,omitempty"`
}

func (o RobotInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RobotInfo struct{}"
	}

	return strings.Join([]string{"RobotInfo", string(data)}, " ")
}

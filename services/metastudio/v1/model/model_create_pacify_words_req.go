package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePacifyWordsReq 创建安抚话术请求。
type CreatePacifyWordsReq struct {

	// 应用ID。
	RobotId string `json:"robot_id"`

	// 安抚话术类型 > 0:通用安抚话术, 1:意图匹配安抚话术
	PacifyWordsType int32 `json:"pacify_words_type"`

	// 意图名称
	Intent *string `json:"intent,omitempty"`

	// 安抚话术。
	PacifyWords string `json:"pacify_words"`

	Language *LanguageEnum `json:"language"`
}

func (o CreatePacifyWordsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePacifyWordsReq struct{}"
	}

	return strings.Join([]string{"CreatePacifyWordsReq", string(data)}, " ")
}

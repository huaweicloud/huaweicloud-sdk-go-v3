package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePacifyWordsResponse Response Object
type UpdatePacifyWordsResponse struct {

	// 安抚话术ID。
	PacifyWordsId *string `json:"pacify_words_id,omitempty"`

	// 安抚话术。
	PacifyWords *string `json:"pacify_words,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	// 安抚话术类型 > 0:通用安抚话术, 1:基于意图匹配安抚话术
	PacifyWordsType *int32 `json:"pacify_words_type,omitempty"`

	// 意图名称
	Intent *string `json:"intent,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePacifyWordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePacifyWordsResponse struct{}"
	}

	return strings.Join([]string{"UpdatePacifyWordsResponse", string(data)}, " ")
}

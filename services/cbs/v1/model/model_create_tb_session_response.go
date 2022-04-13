package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTbSessionResponse struct {
	// 问题。

	Question *string `json:"question,omitempty"`
	// 0表示继续， 1表示直接中断， 2表示播放结束音后中断。

	Action *int32 `json:"action,omitempty"`
	// 会话ID。

	SessionId *string `json:"session_id,omitempty"`
	// 问题ID。

	QuestionId *string `json:"question_id,omitempty"`
	// 语音文件地址。

	AudioFilePath  *string `json:"audio_file_path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTbSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTbSessionResponse struct{}"
	}

	return strings.Join([]string{"CreateTbSessionResponse", string(data)}, " ")
}

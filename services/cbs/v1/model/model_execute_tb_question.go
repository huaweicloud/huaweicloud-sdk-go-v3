package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ExecuteTbQuestion struct {
	// 问题ID。

	QuestionId string `json:"question_id"`
	// 语音文件路径。

	AudioFilePath *string `json:"audio_file_path,omitempty"`
	// 问题。

	Question string `json:"question"`
	// 0 继续， 1 直接中断， 2 播放结束音后中断。

	Action int32 `json:"action"`
}

func (o ExecuteTbQuestion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTbQuestion struct{}"
	}

	return strings.Join([]string{"ExecuteTbQuestion", string(data)}, " ")
}

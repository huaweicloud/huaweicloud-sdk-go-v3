package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubtitleFileReq struct {

	// 剧本ID。
	ScriptId *string `json:"script_id,omitempty"`

	// 剧本序号。  > * 剧本序号不填生成所有场景的字幕；如果需要生成单场景的字幕，需要填剧本序号。
	SequenceNo *int32 `json:"sequence_no,omitempty"`

	CallbackConfig *CallBackConfig `json:"callback_config,omitempty"`
}

func (o CreateSubtitleFileReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubtitleFileReq struct{}"
	}

	return strings.Join([]string{"CreateSubtitleFileReq", string(data)}, " ")
}

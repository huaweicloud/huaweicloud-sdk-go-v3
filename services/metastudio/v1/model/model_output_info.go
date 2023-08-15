package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutputInfo 输出信息。
type OutputInfo struct {

	// 面部表情输入地址。
	FaceAddr *string `json:"face_addr,omitempty"`

	// 身体动作输入地址。
	BodyAddr *string `json:"body_addr,omitempty"`

	// 音频输入地址。
	AudioAddr *string `json:"audio_addr,omitempty"`

	// 会话ID。
	SessionId *int32 `json:"session_id,omitempty"`
}

func (o OutputInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputInfo struct{}"
	}

	return strings.Join([]string{"OutputInfo", string(data)}, " ")
}

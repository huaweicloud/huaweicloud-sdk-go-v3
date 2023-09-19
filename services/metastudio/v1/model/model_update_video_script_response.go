package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVideoScriptResponse Response Object
type UpdateVideoScriptResponse struct {

	// 剧本ID
	ScriptId *string `json:"script_id,omitempty"`

	AudioFiles *ShootScriptAudioFiles `json:"audio_files,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVideoScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVideoScriptResponse struct{}"
	}

	return strings.Join([]string{"UpdateVideoScriptResponse", string(data)}, " ")
}

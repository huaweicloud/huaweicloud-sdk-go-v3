package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVideoScriptsResponse Response Object
type CreateVideoScriptsResponse struct {

	// 剧本ID
	ScriptId *string `json:"script_id,omitempty"`

	AudioFiles *ShootScriptAudioFiles `json:"audio_files,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVideoScriptsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoScriptsResponse struct{}"
	}

	return strings.Join([]string{"CreateVideoScriptsResponse", string(data)}, " ")
}

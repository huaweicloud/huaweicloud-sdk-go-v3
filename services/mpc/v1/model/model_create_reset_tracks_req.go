package model

import (
	"encoding/json"

	"strings"
)

type CreateResetTracksReq struct {
	AudioFile *SubAudioFile `json:"audio_file,omitempty"`
}

func (o CreateResetTracksReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateResetTracksReq struct{}"
	}

	return strings.Join([]string{"CreateResetTracksReq", string(data)}, " ")
}

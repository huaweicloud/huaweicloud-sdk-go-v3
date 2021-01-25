package model

import (
	"encoding/json"

	"strings"
)

type CreateMergeChannelsReq struct {
	MultiAudio *MpcMultiAudio `json:"multi_audio,omitempty"`
}

func (o CreateMergeChannelsReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMergeChannelsReq struct{}"
	}

	return strings.Join([]string{"CreateMergeChannelsReq", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

type ExtractAudioTaskReq struct {
	// 媒资id<br/>

	AssetId string `json:"asset_id"`

	Parameter *Parameter `json:"parameter,omitempty"`
}

func (o ExtractAudioTaskReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExtractAudioTaskReq struct{}"
	}

	return strings.Join([]string{"ExtractAudioTaskReq", string(data)}, " ")
}

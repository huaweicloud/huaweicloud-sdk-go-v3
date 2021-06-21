package model

import (
	"encoding/json"

	"strings"
)

type DefaultRecordConfig struct {
	// 录制格式flv，hls，mp4

	RecordFormat []VideoFormatVar `json:"record_format"`

	ObsAddr *RecordObsFileAddr `json:"obs_addr"`

	HlsConfig *HlsRecordConfig `json:"hls_config,omitempty"`

	FlvConfig *FlvRecordConfig `json:"flv_config,omitempty"`

	Mp4Config *Mp4RecordConfig `json:"mp4_config,omitempty"`
}

func (o DefaultRecordConfig) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DefaultRecordConfig struct{}"
	}

	return strings.Join([]string{"DefaultRecordConfig", string(data)}, " ")
}

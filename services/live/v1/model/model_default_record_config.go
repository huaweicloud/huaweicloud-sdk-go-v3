package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DefaultRecordConfig struct {

	// 录制格式，当前支持：FLV，HLS，MP4三种格式，设置格式时必须使用大写字母
	RecordFormat []VideoFormatVar `json:"record_format" xml:"record_format"`

	ObsAddr *RecordObsFileAddr `json:"obs_addr" xml:"obs_addr"`

	HlsConfig *HlsRecordConfig `json:"hls_config,omitempty" xml:"hls_config"`

	FlvConfig *FlvRecordConfig `json:"flv_config,omitempty" xml:"flv_config"`

	Mp4Config *Mp4RecordConfig `json:"mp4_config,omitempty" xml:"mp4_config"`
}

func (o DefaultRecordConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefaultRecordConfig struct{}"
	}

	return strings.Join([]string{"DefaultRecordConfig", string(data)}, " ")
}

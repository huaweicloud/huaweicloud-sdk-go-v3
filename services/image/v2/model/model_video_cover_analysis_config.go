package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoCoverAnalysisConfig struct {
	Common *VideoCoverAnalysisConfigCommon `json:"common,omitempty"`
}

func (o VideoCoverAnalysisConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCoverAnalysisConfig struct{}"
	}

	return strings.Join([]string{"VideoCoverAnalysisConfig", string(data)}, " ")
}

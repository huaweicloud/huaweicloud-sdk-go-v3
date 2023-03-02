package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoCuttingConfig struct {
	Common *VideoCuttingConfigCommon `json:"common,omitempty"`
}

func (o VideoCuttingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCuttingConfig struct{}"
	}

	return strings.Join([]string{"VideoCuttingConfig", string(data)}, " ")
}

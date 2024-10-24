package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogConfig Ray日志配置
type LogConfig struct {

	// 是否开启日志并记录到LTS,默认不开启
	EnableLtsLog *bool `json:"enable_lts_log,omitempty"`

	// 是否开启日志转储到OBS，默认不开启 开启的时候需要先开启日志记录到LTS，否则无法开启。
	EnableObsLog *bool `json:"enable_obs_log,omitempty"`

	// 日志存储的OBS路径
	ObsPath *string `json:"obs_path,omitempty"`
}

func (o LogConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogConfig struct{}"
	}

	return strings.Join([]string{"LogConfig", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoteObs 数据实际输出到OBS。
type RemoteObs struct {

	// 数据实际输出到OBS的路径。
	ObsUrl string `json:"obs_url"`
}

func (o RemoteObs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteObs struct{}"
	}

	return strings.Join([]string{"RemoteObs", string(data)}, " ")
}

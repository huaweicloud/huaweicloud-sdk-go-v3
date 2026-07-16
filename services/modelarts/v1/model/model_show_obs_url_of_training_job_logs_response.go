package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObsUrlOfTrainingJobLogsResponse Response Object
type ShowObsUrlOfTrainingJobLogsResponse struct {

	// 日志OBS临时链接（复制到浏览器可查看当前全量日志）。
	ObsUrl *string `json:"obs_url,omitempty"`

	Shards         *Shards `json:"shards,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowObsUrlOfTrainingJobLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObsUrlOfTrainingJobLogsResponse struct{}"
	}

	return strings.Join([]string{"ShowObsUrlOfTrainingJobLogsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartPipelineRequest struct {
	// 指定待开启的集群ID。

	ClusterId string `json:"cluster_id"`

	Body *StartPipelineReq `json:"body,omitempty"`
}

func (o StartPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineRequest struct{}"
	}

	return strings.Join([]string{"StartPipelineRequest", string(data)}, " ")
}

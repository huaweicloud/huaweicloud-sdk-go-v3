package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartHotPipelineRequest Request Object
type StartHotPipelineRequest struct {

	// 指定待操作的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *StartHotPipelineRequestBody `json:"body,omitempty"`
}

func (o StartHotPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartHotPipelineRequest struct{}"
	}

	return strings.Join([]string{"StartHotPipelineRequest", string(data)}, " ")
}

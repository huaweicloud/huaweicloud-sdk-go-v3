package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartPipelineJobResponse Response Object
type StartPipelineJobResponse struct {

	// 管道ID
	PipelineId     *string `json:"pipeline_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartPipelineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineJobResponse struct{}"
	}

	return strings.Join([]string{"StartPipelineJobResponse", string(data)}, " ")
}

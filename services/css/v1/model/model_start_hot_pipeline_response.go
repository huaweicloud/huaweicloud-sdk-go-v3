package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartHotPipelineResponse Response Object
type StartHotPipelineResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartHotPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartHotPipelineResponse struct{}"
	}

	return strings.Join([]string{"StartHotPipelineResponse", string(data)}, " ")
}

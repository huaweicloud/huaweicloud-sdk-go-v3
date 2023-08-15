package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartPipelineResponse Response Object
type StartPipelineResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineResponse struct{}"
	}

	return strings.Join([]string{"StartPipelineResponse", string(data)}, " ")
}

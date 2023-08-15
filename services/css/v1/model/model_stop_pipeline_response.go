package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopPipelineResponse Response Object
type StopPipelineResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPipelineResponse struct{}"
	}

	return strings.Join([]string{"StopPipelineResponse", string(data)}, " ")
}

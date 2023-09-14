package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowPipelinesStatusResponse Response Object
type BatchShowPipelinesStatusResponse struct {
	Body           *[]PipelineExecuteStates `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchShowPipelinesStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowPipelinesStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchShowPipelinesStatusResponse", string(data)}, " ")
}

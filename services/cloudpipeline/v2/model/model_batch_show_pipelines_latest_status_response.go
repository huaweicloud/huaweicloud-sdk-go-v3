package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowPipelinesLatestStatusResponse Response Object
type BatchShowPipelinesLatestStatusResponse struct {
	Body           *[]PipelineLatestRun `json:"body,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchShowPipelinesLatestStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowPipelinesLatestStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchShowPipelinesLatestStatusResponse", string(data)}, " ")
}

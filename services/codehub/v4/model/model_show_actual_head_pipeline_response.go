package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowActualHeadPipelineResponse Response Object
type ShowActualHeadPipelineResponse struct {
	Data *PipelineDetailDto `json:"data,omitempty"`

	// 最新的commit是否有对应的流水线
	IsValid        *bool `json:"is_valid,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowActualHeadPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActualHeadPipelineResponse struct{}"
	}

	return strings.Join([]string{"ShowActualHeadPipelineResponse", string(data)}, " ")
}

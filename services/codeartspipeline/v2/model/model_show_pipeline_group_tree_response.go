package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineGroupTreeResponse Response Object
type ShowPipelineGroupTreeResponse struct {
	Body           *[]PipelineGroupVo `json:"body,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowPipelineGroupTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineGroupTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineGroupTreeResponse", string(data)}, " ")
}

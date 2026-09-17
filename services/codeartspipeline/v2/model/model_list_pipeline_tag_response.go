package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineTagResponse Response Object
type ListPipelineTagResponse struct {
	Body           *[]PipelineTagResp `json:"body,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListPipelineTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineTagResponse struct{}"
	}

	return strings.Join([]string{"ListPipelineTagResponse", string(data)}, " ")
}

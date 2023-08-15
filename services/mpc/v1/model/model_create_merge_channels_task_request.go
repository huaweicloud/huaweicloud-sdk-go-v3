package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMergeChannelsTaskRequest Request Object
type CreateMergeChannelsTaskRequest struct {
	Body *CreateMergeChannelsReq `json:"body,omitempty"`
}

func (o CreateMergeChannelsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeChannelsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateMergeChannelsTaskRequest", string(data)}, " ")
}

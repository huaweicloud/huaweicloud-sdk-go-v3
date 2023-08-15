package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageBatchModerationRequest Request Object
type RunImageBatchModerationRequest struct {
	Body *ImageBatchModerationReq `json:"body,omitempty"`
}

func (o RunImageBatchModerationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageBatchModerationRequest struct{}"
	}

	return strings.Join([]string{"RunImageBatchModerationRequest", string(data)}, " ")
}

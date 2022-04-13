package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVideoShotSplitTaskRequest struct {
	Body *CreateVideoSplitTaskRequestBody `json:"body,omitempty"`
}

func (o CreateVideoShotSplitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoShotSplitTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateVideoShotSplitTaskRequest", string(data)}, " ")
}

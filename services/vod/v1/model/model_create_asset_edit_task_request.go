package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetEditTaskRequest Request Object
type CreateAssetEditTaskRequest struct {
	Body *CreateAssetEditTaskReq `json:"body,omitempty"`
}

func (o CreateAssetEditTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetEditTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetEditTaskRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDevServerJobRequest Request Object
type CreateDevServerJobRequest struct {
	Body *DevServerJobCreateRequest `json:"body,omitempty"`
}

func (o CreateDevServerJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDevServerJobRequest struct{}"
	}

	return strings.Join([]string{"CreateDevServerJobRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateImageToVideoTaskRequest struct {
	Body *ImageToVideoRequestBody `json:"body,omitempty"`
}

func (o CreateImageToVideoTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageToVideoTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateImageToVideoTaskRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStreamJobRequest Request Object
type CreateStreamJobRequest struct {
	Body *CreateStreamJobRequestBody `json:"body,omitempty"`
}

func (o CreateStreamJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamJobRequest struct{}"
	}

	return strings.Join([]string{"CreateStreamJobRequest", string(data)}, " ")
}

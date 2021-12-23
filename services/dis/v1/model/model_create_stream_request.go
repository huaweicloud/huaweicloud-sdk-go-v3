package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateStreamRequest struct {
	Body *CreateStreamReq `json:"body,omitempty"`
}

func (o CreateStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamRequest struct{}"
	}

	return strings.Join([]string{"CreateStreamRequest", string(data)}, " ")
}

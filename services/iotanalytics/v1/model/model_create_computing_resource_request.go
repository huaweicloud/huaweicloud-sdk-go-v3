package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateComputingResourceRequest struct {
	Body *CreateComputingResourceRequestBody `json:"body,omitempty"`
}

func (o CreateComputingResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComputingResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateComputingResourceRequest", string(data)}, " ")
}

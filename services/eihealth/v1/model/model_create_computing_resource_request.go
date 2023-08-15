package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComputingResourceRequest Request Object
type CreateComputingResourceRequest struct {
	Body *CreateComputingResourceReq `json:"body,omitempty"`
}

func (o CreateComputingResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComputingResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateComputingResourceRequest", string(data)}, " ")
}

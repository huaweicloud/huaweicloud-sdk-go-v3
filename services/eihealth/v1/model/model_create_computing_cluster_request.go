package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComputingClusterRequest Request Object
type CreateComputingClusterRequest struct {
	Body *CreateComputingClusterReq `json:"body,omitempty"`
}

func (o CreateComputingClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComputingClusterRequest struct{}"
	}

	return strings.Join([]string{"CreateComputingClusterRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHyperClusterRequest Request Object
type CreateHyperClusterRequest struct {
	Body *HyperClusterCreateRequest `json:"body,omitempty"`
}

func (o CreateHyperClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHyperClusterRequest struct{}"
	}

	return strings.Join([]string{"CreateHyperClusterRequest", string(data)}, " ")
}

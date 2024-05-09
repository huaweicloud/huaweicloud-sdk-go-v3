package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCssClusterRequest Request Object
type CreateCssClusterRequest struct {
	Body *CreateCssClusterReq `json:"body,omitempty"`
}

func (o CreateCssClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCssClusterRequest struct{}"
	}

	return strings.Join([]string{"CreateCssClusterRequest", string(data)}, " ")
}

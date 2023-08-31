package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHostClusterRequest Request Object
type CreateHostClusterRequest struct {
	Body *CreateHostClusterRequestBody `json:"body,omitempty"`
}

func (o CreateHostClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostClusterRequest struct{}"
	}

	return strings.Join([]string{"CreateHostClusterRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEvolveTaskRequest Request Object
type CreateEvolveTaskRequest struct {
	Body *EvolveTaskCreateReq `json:"body,omitempty"`
}

func (o CreateEvolveTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEvolveTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateEvolveTaskRequest", string(data)}, " ")
}

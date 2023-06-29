package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEvaluationProjectRequest Request Object
type CreateEvaluationProjectRequest struct {
	Body *CreateEvaluationProjectReq `json:"body,omitempty"`
}

func (o CreateEvaluationProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEvaluationProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateEvaluationProjectRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEvaluationProjectResponse Response Object
type CreateEvaluationProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateEvaluationProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEvaluationProjectResponse struct{}"
	}

	return strings.Join([]string{"CreateEvaluationProjectResponse", string(data)}, " ")
}

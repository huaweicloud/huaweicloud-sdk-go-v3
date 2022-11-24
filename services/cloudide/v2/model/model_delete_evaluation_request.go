package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEvaluationRequest struct {

	// 评论id
	EvaluationId int64 `json:"evaluation_id"`
}

func (o DeleteEvaluationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEvaluationRequest struct{}"
	}

	return strings.Join([]string{"DeleteEvaluationRequest", string(data)}, " ")
}

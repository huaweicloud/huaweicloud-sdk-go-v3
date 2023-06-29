package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEvaluationResponse Response Object
type DeleteEvaluationResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEvaluationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEvaluationResponse struct{}"
	}

	return strings.Join([]string{"DeleteEvaluationResponse", string(data)}, " ")
}

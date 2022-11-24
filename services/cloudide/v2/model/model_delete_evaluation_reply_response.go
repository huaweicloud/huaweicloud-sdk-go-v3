package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEvaluationReplyResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEvaluationReplyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEvaluationReplyResponse struct{}"
	}

	return strings.Join([]string{"DeleteEvaluationReplyResponse", string(data)}, " ")
}

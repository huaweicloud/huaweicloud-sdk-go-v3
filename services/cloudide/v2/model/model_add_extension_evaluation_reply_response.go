package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddExtensionEvaluationReplyResponse Response Object
type AddExtensionEvaluationReplyResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddExtensionEvaluationReplyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddExtensionEvaluationReplyResponse struct{}"
	}

	return strings.Join([]string{"AddExtensionEvaluationReplyResponse", string(data)}, " ")
}

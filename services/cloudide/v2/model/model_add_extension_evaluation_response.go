package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddExtensionEvaluationResponse Response Object
type AddExtensionEvaluationResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddExtensionEvaluationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddExtensionEvaluationResponse struct{}"
	}

	return strings.Join([]string{"AddExtensionEvaluationResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtensionEvaluationResponse Response Object
type ShowExtensionEvaluationResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowExtensionEvaluationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtensionEvaluationResponse struct{}"
	}

	return strings.Join([]string{"ShowExtensionEvaluationResponse", string(data)}, " ")
}

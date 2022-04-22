package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowHasPipelineResponse struct {
	Error *Error `json:"error,omitempty"`

	// 响应结果
	Result *bool `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHasPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHasPipelineResponse struct{}"
	}

	return strings.Join([]string{"ShowHasPipelineResponse", string(data)}, " ")
}

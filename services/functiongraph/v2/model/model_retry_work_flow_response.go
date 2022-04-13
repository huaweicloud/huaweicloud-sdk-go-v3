package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RetryWorkFlowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RetryWorkFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryWorkFlowResponse struct{}"
	}

	return strings.Join([]string{"RetryWorkFlowResponse", string(data)}, " ")
}

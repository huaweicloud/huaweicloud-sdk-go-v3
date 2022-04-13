package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopWorkFlowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopWorkFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopWorkFlowResponse struct{}"
	}

	return strings.Join([]string{"StopWorkFlowResponse", string(data)}, " ")
}

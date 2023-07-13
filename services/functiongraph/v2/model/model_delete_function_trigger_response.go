package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionTriggerResponse Response Object
type DeleteFunctionTriggerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFunctionTriggerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionTriggerResponse struct{}"
	}

	return strings.Join([]string{"DeleteFunctionTriggerResponse", string(data)}, " ")
}

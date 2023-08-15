package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTriggerResponse Response Object
type DeleteTriggerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTriggerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTriggerResponse struct{}"
	}

	return strings.Join([]string{"DeleteTriggerResponse", string(data)}, " ")
}

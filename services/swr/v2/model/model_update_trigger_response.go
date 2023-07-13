package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTriggerResponse Response Object
type UpdateTriggerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTriggerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTriggerResponse struct{}"
	}

	return strings.Join([]string{"UpdateTriggerResponse", string(data)}, " ")
}

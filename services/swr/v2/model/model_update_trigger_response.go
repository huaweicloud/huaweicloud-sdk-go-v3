package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTriggerResponse Response Object
type CreateTriggerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTriggerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTriggerResponse struct{}"
	}

	return strings.Join([]string{"CreateTriggerResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduledTaskResponse Response Object
type UpdateScheduledTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateScheduledTaskResponse", string(data)}, " ")
}

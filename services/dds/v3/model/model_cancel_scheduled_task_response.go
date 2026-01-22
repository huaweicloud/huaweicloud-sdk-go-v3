package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelScheduledTaskResponse Response Object
type CancelScheduledTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelScheduledTaskResponse", string(data)}, " ")
}

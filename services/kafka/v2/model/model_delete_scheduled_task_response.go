package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduledTaskResponse Response Object
type DeleteScheduledTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteScheduledTaskResponse", string(data)}, " ")
}

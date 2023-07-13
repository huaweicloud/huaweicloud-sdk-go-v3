package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupScheduledTaskResponse Response Object
type DeleteGroupScheduledTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGroupScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteGroupScheduledTaskResponse", string(data)}, " ")
}

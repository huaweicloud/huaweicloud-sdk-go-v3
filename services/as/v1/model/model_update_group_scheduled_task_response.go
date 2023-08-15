package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupScheduledTaskResponse Response Object
type UpdateGroupScheduledTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateGroupScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateGroupScheduledTaskResponse", string(data)}, " ")
}

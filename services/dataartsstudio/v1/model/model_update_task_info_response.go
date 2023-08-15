package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskInfoResponse Response Object
type UpdateTaskInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTaskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskInfoResponse", string(data)}, " ")
}

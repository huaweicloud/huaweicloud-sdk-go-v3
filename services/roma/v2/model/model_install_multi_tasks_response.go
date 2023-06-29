package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallMultiTasksResponse Response Object
type InstallMultiTasksResponse struct {
	Body           *[]MultiTaskInitElement `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o InstallMultiTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallMultiTasksResponse struct{}"
	}

	return strings.Join([]string{"InstallMultiTasksResponse", string(data)}, " ")
}

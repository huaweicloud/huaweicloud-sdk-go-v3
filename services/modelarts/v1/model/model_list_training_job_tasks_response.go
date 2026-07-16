package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobTasksResponse Response Object
type ListTrainingJobTasksResponse struct {
	Body           *[]TaskHistory `json:"body,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListTrainingJobTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTrainingJobTasksResponse", string(data)}, " ")
}

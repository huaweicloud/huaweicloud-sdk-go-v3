package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTasksResponse Response Object
type CreateTasksResponse struct {
	Body           *[]CreateResponseBody `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CreateTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTasksResponse struct{}"
	}

	return strings.Join([]string{"CreateTasksResponse", string(data)}, " ")
}

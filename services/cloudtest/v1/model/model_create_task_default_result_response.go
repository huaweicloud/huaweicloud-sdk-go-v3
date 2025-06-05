package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskDefaultResultResponse Response Object
type CreateTaskDefaultResultResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTaskDefaultResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskDefaultResultResponse struct{}"
	}

	return strings.Join([]string{"CreateTaskDefaultResultResponse", string(data)}, " ")
}

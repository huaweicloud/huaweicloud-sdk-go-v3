package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCpiTaskResponse Response Object
type CreateCpiTaskResponse struct {

	// CPI任务ID
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCpiTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCpiTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateCpiTaskResponse", string(data)}, " ")
}

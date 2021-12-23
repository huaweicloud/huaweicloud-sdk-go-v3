package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEditingJobResponse struct {
	Body           *[]QueryEditingJobRsp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListEditingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEditingJobResponse struct{}"
	}

	return strings.Join([]string{"ListEditingJobResponse", string(data)}, " ")
}

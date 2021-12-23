package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTransferTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTransferTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateTransferTaskResponse", string(data)}, " ")
}

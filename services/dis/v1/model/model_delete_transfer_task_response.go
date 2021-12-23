package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTransferTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTransferTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransferTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteTransferTaskResponse", string(data)}, " ")
}

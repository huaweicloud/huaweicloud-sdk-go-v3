package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferProjectResponse Response Object
type TransferProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o TransferProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferProjectResponse struct{}"
	}

	return strings.Join([]string{"TransferProjectResponse", string(data)}, " ")
}

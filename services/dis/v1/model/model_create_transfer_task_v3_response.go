package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTransferTaskV3Response struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTransferTaskV3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferTaskV3Response struct{}"
	}

	return strings.Join([]string{"CreateTransferTaskV3Response", string(data)}, " ")
}

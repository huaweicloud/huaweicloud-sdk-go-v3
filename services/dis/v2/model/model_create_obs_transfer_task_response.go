package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObsTransferTaskResponse Response Object
type CreateObsTransferTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateObsTransferTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObsTransferTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateObsTransferTaskResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferFilePreRequest Request Object
type TransferFilePreRequest struct {
	Body *TransferFilePreReq `json:"body,omitempty"`
}

func (o TransferFilePreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferFilePreRequest struct{}"
	}

	return strings.Join([]string{"TransferFilePreRequest", string(data)}, " ")
}

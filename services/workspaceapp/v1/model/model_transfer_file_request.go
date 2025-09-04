package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferFileRequest Request Object
type TransferFileRequest struct {
	Body *TransferFileReq `json:"body,omitempty"`
}

func (o TransferFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferFileRequest struct{}"
	}

	return strings.Join([]string{"TransferFileRequest", string(data)}, " ")
}

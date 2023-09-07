package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportOperationInfoResponse Response Object
type ExportOperationInfoResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportOperationInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportOperationInfoResponse struct{}"
	}

	return strings.Join([]string{"ExportOperationInfoResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSessionsRequest Request Object
type ExportSessionsRequest struct {
	Body *ExportSessionsReq `json:"body,omitempty"`
}

func (o ExportSessionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSessionsRequest struct{}"
	}

	return strings.Join([]string{"ExportSessionsRequest", string(data)}, " ")
}

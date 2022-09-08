package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportConnectionsRequest struct {
	Body *ImportConnectionReq `json:"body,omitempty"`
}

func (o ImportConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ImportConnectionsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserConnectionRequest Request Object
type ExportUserConnectionRequest struct {
	Body *ListUserConnectionReq `json:"body,omitempty"`
}

func (o ExportUserConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserConnectionRequest struct{}"
	}

	return strings.Join([]string{"ExportUserConnectionRequest", string(data)}, " ")
}

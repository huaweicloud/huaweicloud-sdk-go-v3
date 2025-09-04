package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAppConnectionRequest Request Object
type ExportAppConnectionRequest struct {
	Body *ListAppConnectionReq `json:"body,omitempty"`
}

func (o ExportAppConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAppConnectionRequest struct{}"
	}

	return strings.Join([]string{"ExportAppConnectionRequest", string(data)}, " ")
}

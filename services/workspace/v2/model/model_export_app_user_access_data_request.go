package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAppUserAccessDataRequest Request Object
type ExportAppUserAccessDataRequest struct {
	Body *ExportAppUserAccessDataRequestBody `json:"body,omitempty"`
}

func (o ExportAppUserAccessDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAppUserAccessDataRequest struct{}"
	}

	return strings.Join([]string{"ExportAppUserAccessDataRequest", string(data)}, " ")
}

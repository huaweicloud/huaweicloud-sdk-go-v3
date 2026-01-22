package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExportStatusResponse Response Object
type ShowExportStatusResponse struct {
	Data           *ExportStatusVo `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowExportStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExportStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowExportStatusResponse", string(data)}, " ")
}

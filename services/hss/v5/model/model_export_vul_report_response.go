package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportVulReportResponse Response Object
type ExportVulReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportVulReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVulReportResponse struct{}"
	}

	return strings.Join([]string{"ExportVulReportResponse", string(data)}, " ")
}

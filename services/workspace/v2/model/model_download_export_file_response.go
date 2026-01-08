package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadExportFileResponse Response Object
type DownloadExportFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadExportFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadExportFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadExportFileResponse", string(data)}, " ")
}

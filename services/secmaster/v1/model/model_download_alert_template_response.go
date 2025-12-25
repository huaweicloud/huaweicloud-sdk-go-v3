package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAlertTemplateResponse Response Object
type DownloadAlertTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadAlertTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAlertTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadAlertTemplateResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadIncidentTemplateResponse Response Object
type DownloadIncidentTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadIncidentTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadIncidentTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadIncidentTemplateResponse", string(data)}, " ")
}

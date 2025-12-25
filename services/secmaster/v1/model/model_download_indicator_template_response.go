package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadIndicatorTemplateResponse Response Object
type DownloadIndicatorTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadIndicatorTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadIndicatorTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadIndicatorTemplateResponse", string(data)}, " ")
}

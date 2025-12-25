package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadResourceTemplateResponse Response Object
type DownloadResourceTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadResourceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadResourceTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadResourceTemplateResponse", string(data)}, " ")
}

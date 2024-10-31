package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAssetTemplateResponse Response Object
type DownloadAssetTemplateResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadAssetTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAssetTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadAssetTemplateResponse", string(data)}, " ")
}

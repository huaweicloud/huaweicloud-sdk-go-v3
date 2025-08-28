package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAssetFileResponse Response Object
type DownloadAssetFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadAssetFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAssetFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadAssetFileResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadMetadataResponse Response Object
type DownloadMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadMetadataResponse struct{}"
	}

	return strings.Join([]string{"DownloadMetadataResponse", string(data)}, " ")
}

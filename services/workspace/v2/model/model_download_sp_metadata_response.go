package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadSpMetadataResponse Response Object
type DownloadSpMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadSpMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSpMetadataResponse struct{}"
	}

	return strings.Join([]string{"DownloadSpMetadataResponse", string(data)}, " ")
}

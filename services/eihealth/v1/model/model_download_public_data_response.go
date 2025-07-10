package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadPublicDataResponse Response Object
type DownloadPublicDataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadPublicDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadPublicDataResponse struct{}"
	}

	return strings.Join([]string{"DownloadPublicDataResponse", string(data)}, " ")
}

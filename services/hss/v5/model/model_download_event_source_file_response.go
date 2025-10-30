package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadEventSourceFileResponse Response Object
type DownloadEventSourceFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadEventSourceFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadEventSourceFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadEventSourceFileResponse", string(data)}, " ")
}

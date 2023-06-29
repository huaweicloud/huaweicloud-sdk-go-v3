package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadDataJobLogResponse Response Object
type DownloadDataJobLogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadDataJobLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDataJobLogResponse struct{}"
	}

	return strings.Join([]string{"DownloadDataJobLogResponse", string(data)}, " ")
}

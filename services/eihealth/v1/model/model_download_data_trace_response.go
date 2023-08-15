package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadDataTraceResponse Response Object
type DownloadDataTraceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadDataTraceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDataTraceResponse struct{}"
	}

	return strings.Join([]string{"DownloadDataTraceResponse", string(data)}, " ")
}

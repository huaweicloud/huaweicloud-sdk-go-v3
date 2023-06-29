package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAppVersionResponse Response Object
type DownloadAppVersionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadAppVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAppVersionResponse struct{}"
	}

	return strings.Join([]string{"DownloadAppVersionResponse", string(data)}, " ")
}

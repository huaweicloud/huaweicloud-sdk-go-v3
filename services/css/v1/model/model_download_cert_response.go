package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadCertResponse Response Object
type DownloadCertResponse struct {

	// 文件流
	File           *string `json:"file,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadCertResponse struct{}"
	}

	return strings.Join([]string{"DownloadCertResponse", string(data)}, " ")
}

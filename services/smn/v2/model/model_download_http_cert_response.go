package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadHttpCertResponse Response Object
type DownloadHttpCertResponse struct {
	ContentType    *string `json:"content-type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadHttpCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadHttpCertResponse struct{}"
	}

	return strings.Join([]string{"DownloadHttpCertResponse", string(data)}, " ")
}

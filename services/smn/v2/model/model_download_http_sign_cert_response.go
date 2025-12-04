package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadHttpSignCertResponse Response Object
type DownloadHttpSignCertResponse struct {
	ContentType    *string `json:"content-type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadHttpSignCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadHttpSignCertResponse struct{}"
	}

	return strings.Join([]string{"DownloadHttpSignCertResponse", string(data)}, " ")
}

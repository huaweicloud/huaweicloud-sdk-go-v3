package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSslCertDownloadLinkResponse Response Object
type ShowSslCertDownloadLinkResponse struct {

	// ssl下载链接。
	DownloadLink   *string `json:"download_link,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSslCertDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSslCertDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowSslCertDownloadLinkResponse", string(data)}, " ")
}

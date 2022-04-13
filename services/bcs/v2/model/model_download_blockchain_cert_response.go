package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DownloadBlockchainCertResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadBlockchainCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBlockchainCertResponse struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainCertResponse", string(data)}, " ")
}

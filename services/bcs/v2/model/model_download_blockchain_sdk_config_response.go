package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DownloadBlockchainSdkConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadBlockchainSdkConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBlockchainSdkConfigResponse struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainSdkConfigResponse", string(data)}, " ")
}

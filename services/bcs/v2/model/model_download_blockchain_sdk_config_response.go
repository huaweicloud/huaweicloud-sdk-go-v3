package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DownloadBlockchainSdkConfigResponse struct {

	// 请求成功的结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadBlockchainSdkConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBlockchainSdkConfigResponse struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainSdkConfigResponse", string(data)}, " ")
}

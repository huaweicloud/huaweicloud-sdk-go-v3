package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadFederationKubeconfigRequestBody 下载联邦证书的请求体
type DownloadFederationKubeconfigRequestBody struct {

	// kubeconfig中证书的有效期
	Duration int32 `json:"duration"`
}

func (o DownloadFederationKubeconfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadFederationKubeconfigRequestBody struct{}"
	}

	return strings.Join([]string{"DownloadFederationKubeconfigRequestBody", string(data)}, " ")
}

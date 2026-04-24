package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadSpMetadataRequest Request Object
type DownloadSpMetadataRequest struct {

	// 身份提供者名称。
	IdentityProvider string `json:"identity_provider"`

	// 接入服务器地址。
	AccessServerAddress string `json:"access_server_address"`
}

func (o DownloadSpMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSpMetadataRequest struct{}"
	}

	return strings.Join([]string{"DownloadSpMetadataRequest", string(data)}, " ")
}

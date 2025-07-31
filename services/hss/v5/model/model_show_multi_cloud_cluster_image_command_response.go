package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMultiCloudClusterImageCommandResponse Response Object
type ShowMultiCloudClusterImageCommandResponse struct {

	// 上传镜像指令
	ImageCommand *string `json:"image_command,omitempty"`

	// 安装凭证指令
	SecretCommand *string `json:"secret_command,omitempty"`

	// 镜像下载链接
	ImagesDownloadUrl *string `json:"images_download_url,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ShowMultiCloudClusterImageCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMultiCloudClusterImageCommandResponse struct{}"
	}

	return strings.Join([]string{"ShowMultiCloudClusterImageCommandResponse", string(data)}, " ")
}

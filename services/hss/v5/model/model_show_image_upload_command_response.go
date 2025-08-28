package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageUploadCommandResponse Response Object
type ShowImageUploadCommandResponse struct {

	// **参数解释**： 上传镜像指令 **取值范围**： 字符长度1-1024位
	ImageCommand *string `json:"image_command,omitempty"`

	// **参数解释**： 镜像下载链接 **取值范围**： 字符长度1-256位
	ImagesDownloadUrl *string `json:"images_download_url,omitempty"`

	// **参数解释**： SWR镜像获取指令 **取值范围**： 字符长度1-1024位
	SwrImagePullCommand *string `json:"swr_image_pull_command,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowImageUploadCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageUploadCommandResponse struct{}"
	}

	return strings.Join([]string{"ShowImageUploadCommandResponse", string(data)}, " ")
}

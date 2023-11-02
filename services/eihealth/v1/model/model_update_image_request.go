package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageRequest Request Object
type UpdateImageRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 镜像id
	ImageId string `json:"image_id"`

	Body *UpdateImageReq `json:"body,omitempty"`
}

func (o UpdateImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageRequest struct{}"
	}

	return strings.Join([]string{"UpdateImageRequest", string(data)}, " ")
}

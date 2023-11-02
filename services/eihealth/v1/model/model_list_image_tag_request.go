package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageTagRequest Request Object
type ListImageTagRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 镜像id
	ImageId string `json:"image_id"`
}

func (o ListImageTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageTagRequest struct{}"
	}

	return strings.Join([]string{"ListImageTagRequest", string(data)}, " ")
}

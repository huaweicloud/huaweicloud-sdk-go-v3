package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTagRequest Request Object
type BatchDeleteTagRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 镜像id
	ImageId string `json:"image_id"`

	Body *BatchDeleteTagReq `json:"body,omitempty"`
}

func (o BatchDeleteTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagRequest", string(data)}, " ")
}

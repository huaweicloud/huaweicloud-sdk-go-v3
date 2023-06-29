package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageRequest Request Object
type ListImageRequest struct {

	// 镜像类型
	Type *string `json:"type,omitempty"`

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 镜像名称
	Name *string `json:"name,omitempty"`

	// 是否展示无镜像版本的镜像
	ShowEmpty *bool `json:"show_empty,omitempty"`
}

func (o ListImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageRequest struct{}"
	}

	return strings.Join([]string{"ListImageRequest", string(data)}, " ")
}

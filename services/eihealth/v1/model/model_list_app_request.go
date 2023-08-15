package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppRequest Request Object
type ListAppRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用版本
	Version *string `json:"version,omitempty"`
}

func (o ListAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppRequest struct{}"
	}

	return strings.Join([]string{"ListAppRequest", string(data)}, " ")
}

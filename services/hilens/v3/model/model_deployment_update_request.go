package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentUpdateRequest struct {
	Deployment *DeploymentRequest `json:"deployment,omitempty"`

	// 应用部署描述修改，只修改描述不需要传deployment参数。最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty"`

	// 部署标签
	Tags *[]DeploymentTag `json:"tags,omitempty"`
}

func (o DeploymentUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentUpdateRequest struct{}"
	}

	return strings.Join([]string{"DeploymentUpdateRequest", string(data)}, " ")
}

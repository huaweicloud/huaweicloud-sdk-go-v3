package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部署参数配置
type Deployment struct {
	// 部署名称

	Name string `json:"name"`
	// 部署描述

	Description *string `json:"description,omitempty"`
	// 应用部署来源: 边缘市场(iem) or 自定义(userdefined)

	Source *string `json:"source,omitempty"`
	// 应用部署到指定节点组，与node_ids二选一

	GroupId *string `json:"group_id,omitempty"`
	// 应用部署到指定节点，当前只支持一个边缘节点

	NodeIds []string `json:"node_ids"`
	// 节点属性

	Tags *[]Attributes `json:"tags,omitempty"`

	Deployment *CreateAppsInDeploymentV3 `json:"deployment"`
}

func (o Deployment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Deployment struct{}"
	}

	return strings.Join([]string{"Deployment", string(data)}, " ")
}

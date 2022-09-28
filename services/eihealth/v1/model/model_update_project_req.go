package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新项目请求体
type UpdateProjectReq struct {

	// 项目描述
	Description *string `json:"description,omitempty"`

	Status *ProjectStatus `json:"status,omitempty"`

	// 项目标签
	Tags *[]string `json:"tags,omitempty"`

	// 是否为核心项目标记
	IsCore *bool `json:"is_core,omitempty"`
}

func (o UpdateProjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectReq struct{}"
	}

	return strings.Join([]string{"UpdateProjectReq", string(data)}, " ")
}

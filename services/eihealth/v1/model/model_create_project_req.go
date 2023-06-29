package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectReq 创建项目请求体
type CreateProjectReq struct {

	// 项目描述
	Description *string `json:"description,omitempty"`

	// 项目名称
	Name string `json:"name"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`

	// 标签
	IsCore *bool `json:"is_core,omitempty"`
}

func (o CreateProjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectReq struct{}"
	}

	return strings.Join([]string{"CreateProjectReq", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAppTemplatesResult struct {

	// 模板id
	Id *string `json:"id,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 模板执行运行时
	Runtime *string `json:"runtime,omitempty"`

	// 模板使用场景
	Category *string `json:"category,omitempty"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 模板镜像文件（base64编码）
	Image *string `json:"image,omitempty"`
}

func (o ListAppTemplatesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppTemplatesResult struct{}"
	}

	return strings.Join([]string{"ListAppTemplatesResult", string(data)}, " ")
}

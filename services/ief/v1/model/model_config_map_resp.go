package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置项
type ConfigMapResp struct {

	// 配置项ID
	Id string `json:"id" xml:"id"`

	// 配置项名称
	Name string `json:"name" xml:"name"`

	// 配置项描述
	Description string `json:"description" xml:"description"`

	// 配置项键列表
	Configs map[string]string `json:"configs" xml:"configs"`

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 创建时间
	CreatedAt string `json:"created_at" xml:"created_at"`

	// 更新时间
	UpdatedAt string `json:"updated_at" xml:"updated_at"`
}

func (o ConfigMapResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigMapResp struct{}"
	}

	return strings.Join([]string{"ConfigMapResp", string(data)}, " ")
}

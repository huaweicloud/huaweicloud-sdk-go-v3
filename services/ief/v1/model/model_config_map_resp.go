package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置项
type ConfigMapResp struct {

	// 配置项ID
	Id string `json:"id"`

	// 配置项名称
	Name string `json:"name"`

	// 配置项描述
	Description string `json:"description"`

	// 配置项键列表
	Configs map[string]string `json:"configs"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`
}

func (o ConfigMapResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigMapResp struct{}"
	}

	return strings.Join([]string{"ConfigMapResp", string(data)}, " ")
}

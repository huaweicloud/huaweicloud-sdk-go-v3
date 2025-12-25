package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetsRequestBody 创建和更新操作连接的请求体
type CreateAssetsRequestBody struct {

	// 操作连接的名称
	Name string `json:"name"`

	// 操作连接的描述信息
	Description *string `json:"description,omitempty"`

	// 操作连接所属的插件id
	ComponentId string `json:"component_id"`

	// 操作连接所属的插件版本id
	ComponentVersionId string `json:"component_version_id"`

	// 具体的操作连接配置字符串，根据插件的操作连接schema配置对应字段值
	Config string `json:"config"`
}

func (o CreateAssetsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAssetsRequestBody", string(data)}, " ")
}

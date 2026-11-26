package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceArtifactManifestRequest Request Object
type ShowInstanceArtifactManifestRequest struct {

	// 注册表实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 制品仓库名称, 如果您的镜像名称repository_name参数中有/，请先把 / 替换成 %2F 再进行请求。注意：使用curl命令请求，需要把 / 替换成 %252F(即 %2F 的 % 被编码为 %25)再次请求。
	RepositoryName string `json:"repository_name"`

	// 制品摘要
	Reference string `json:"reference"`
}

func (o ShowInstanceArtifactManifestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceArtifactManifestRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceArtifactManifestRequest", string(data)}, " ")
}

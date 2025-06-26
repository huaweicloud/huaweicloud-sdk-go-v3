package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceArtifactRequest Request Object
type ShowInstanceArtifactRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 制品名称
	RepositoryName string `json:"repository_name"`

	// 制品摘要
	Reference string `json:"reference"`
}

func (o ShowInstanceArtifactRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceArtifactRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceArtifactRequest", string(data)}, " ")
}

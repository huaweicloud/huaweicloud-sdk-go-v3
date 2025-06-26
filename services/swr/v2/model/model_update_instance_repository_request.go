package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRepositoryRequest Request Object
type UpdateInstanceRepositoryRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 制品仓库名称
	RepositoryName string `json:"repository_name"`

	Body *UpdateRepoConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRepositoryRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRepositoryRequest", string(data)}, " ")
}

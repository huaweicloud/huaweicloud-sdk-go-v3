package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceRepositoryRequest Request Object
type ShowInstanceRepositoryRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 仓库名称
	RepositoryName string `json:"repository_name"`
}

func (o ShowInstanceRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceRepositoryRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceRepositoryRequest", string(data)}, " ")
}

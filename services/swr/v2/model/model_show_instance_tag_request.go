package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTagRequest Request Object
type ShowInstanceTagRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 仓库名称
	RepositoryName string `json:"repository_name"`

	// Tag名称
	TagName string `json:"tag_name"`
}

func (o ShowInstanceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTagRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceTagRequest", string(data)}, " ")
}

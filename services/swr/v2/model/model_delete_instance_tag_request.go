package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceTagRequest Request Object
type DeleteInstanceTagRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 仓库名称
	RepositoryName string `json:"repository_name"`

	// Tag名称
	TagName string `json:"tag_name"`
}

func (o DeleteInstanceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceTagRequest", string(data)}, " ")
}

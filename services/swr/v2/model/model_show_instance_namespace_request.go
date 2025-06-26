package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceNamespaceRequest Request Object
type ShowInstanceNamespaceRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称，小写字母或数字开头，后面跟小写字母、数字、点、下划线或中划线（其中点、下划线、中划线不能直接连续），小写字母或数字结尾，1-64个字符。
	NamespaceName string `json:"namespace_name"`
}

func (o ShowInstanceNamespaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceNamespaceRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceNamespaceRequest", string(data)}, " ")
}

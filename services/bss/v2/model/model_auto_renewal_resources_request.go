package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AutoRenewalResourcesRequest struct {
	// |参数名称：资源实例ID。您可以调用“2.1-查询客户包年包月资源列表”接口获取资源ID。在设置弹性云服务器自动续费时，能够自动将其挂载的硬盘一并设置为自动续费。| |参数的约束及描述：|

	ResourceId string `json:"resource_id"`
}

func (o AutoRenewalResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoRenewalResourcesRequest struct{}"
	}

	return strings.Join([]string{"AutoRenewalResourcesRequest", string(data)}, " ")
}

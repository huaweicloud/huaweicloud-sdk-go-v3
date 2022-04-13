package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelAutoRenewalResourcesRequest struct {
	// |参数名称：资源实例ID。您可以调用“查询客户包年/包月资源列表”接口获取资源ID。在取消弹性云服务器自动续费的时候，能够自动将其挂载的硬盘一并取消自动续费。| |参数的约束及描述：|

	ResourceId string `json:"resource_id"`
}

func (o CancelAutoRenewalResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAutoRenewalResourcesRequest struct{}"
	}

	return strings.Join([]string{"CancelAutoRenewalResourcesRequest", string(data)}, " ")
}

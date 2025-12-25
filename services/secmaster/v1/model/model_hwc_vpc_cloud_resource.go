package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcVpcCloudResource VPC关联资产类型和数量 取值范围：目前只返回VPC关联的routetable和virsubnet。virsubnet数量为ipv4和ipv6子网总数。
type HwcVpcCloudResource struct {

	// 资产类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资产数量
	ResourceCount *int32 `json:"resource_count,omitempty"`
}

func (o HwcVpcCloudResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcVpcCloudResource struct{}"
	}

	return strings.Join([]string{"HwcVpcCloudResource", string(data)}, " ")
}

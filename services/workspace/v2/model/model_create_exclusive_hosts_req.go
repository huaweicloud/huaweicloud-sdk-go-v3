package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExclusiveHostsReq 创建专享主机请求体。
type CreateExclusiveHostsReq struct {

	// 可用分区。
	AvailabilityZone string `json:"availability_zone"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 追加桌面数量。
	ApplyDesktopQuantity *int32 `json:"apply_desktop_quantity,omitempty"`

	// 购买数量。
	Quantity int32 `json:"quantity"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 产品套餐ID。
	ProductId string `json:"product_id"`

	// 镜像盘列表。
	ImageVolumes []Volume `json:"image_volumes"`

	// 企业项目ID，默认\"0\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 存储盘列表。
	MemoryVolumes []Volume `json:"memory_volumes"`

	// vpc id。
	VpcId string `json:"vpc_id"`

	// 子网id。
	SubnetId string `json:"subnet_id"`

	ResizeExclusiveLites *ResizeExclusiveLitesReq `json:"resize_exclusive_lites,omitempty"`
}

func (o CreateExclusiveHostsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExclusiveHostsReq struct{}"
	}

	return strings.Join([]string{"CreateExclusiveHostsReq", string(data)}, " ")
}

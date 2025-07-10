package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChangeOrderReq 创建变更订单请求体。
type CreateChangeOrderReq struct {

	// 类型 resizeDesktops(变更规格)、expandVolumes(扩容磁盘)。
	Type string `json:"type"`

	// 周期类型，2：包月；3：包年。
	PeriodType *int32 `json:"period_type,omitempty"`

	// 周期数。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 变更后规格产品ID，当是resizeDesktops，必传。
	ResizeProductId *string `json:"resize_product_id,omitempty"`

	// 扩容的云硬盘的ID，当是expandVolumes，必传。
	ExpandVolumeId *string `json:"expand_volume_id,omitempty"`

	// 扩容后云硬盘的总大小，当是expandVolumes，必传。范围10-32760，大小为10的倍数。
	ExpandNewSize *int32 `json:"expand_new_size,omitempty"`

	// 专享主机变更桌面数后桌面路数的总大小，当是jobType是resizeExclusiveLites，必传。
	NewQuantity *int32 `json:"new_quantity,omitempty"`

	// 专享主机桌面数的productId，当是resizeExclusiveLites，必传。
	ExclusiveLitesProductId *string `json:"exclusive_lites_product_id,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ExtendParam *OrderExtendParam `json:"extend_param,omitempty"`
}

func (o CreateChangeOrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChangeOrderReq struct{}"
	}

	return strings.Join([]string{"CreateChangeOrderReq", string(data)}, " ")
}

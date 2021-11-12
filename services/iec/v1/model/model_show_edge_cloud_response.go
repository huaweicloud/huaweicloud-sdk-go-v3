package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEdgeCloudResponse struct {
	// 边缘业务ID。

	Id *string `json:"id,omitempty"`
	// 边缘业务名称。

	Name *string `json:"name,omitempty"`
	// 边缘业务资源组。

	Stacks *[]Stack `json:"stacks,omitempty"`

	Coverage *CoverageResp `json:"coverage,omitempty"`
	// 边缘业务成功创建的虚拟机数量。

	SuccessNum *int32 `json:"success_num,omitempty"`
	// 边缘业务创建失败的虚拟机数量。

	FailedNum *int32 `json:"failed_num,omitempty"`
	// 边缘业务状态。

	Status *string `json:"status,omitempty"`

	FailReason *FailReason `json:"fail_reason,omitempty"`
	// 边缘业务支持的边缘区域数目，等同于边缘业务下所有实例的区域数目总和

	EdgeRegions *int32 `json:"edge_regions,omitempty"`
	// 描述。

	Description *string `json:"description,omitempty"`
	// 创建时间。

	CreateAt *string `json:"create_at,omitempty"`
	// 修改时间。

	UpdateAt       *string `json:"update_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEdgeCloudResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeCloudResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeCloudResponse", string(data)}, " ")
}

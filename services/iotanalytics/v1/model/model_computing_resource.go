package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputingResource struct {
	// 计算资源ID。

	ComputingResourceId *string `json:"computing_resource_id,omitempty"`
	// 计算资源名称。

	ComputingResourceName *string `json:"computing_resource_name,omitempty"`
	// 计算资源描述信息。

	Description *string `json:"description,omitempty"`
	// 创建计算资源的用户。

	Owner *string `json:"owner,omitempty"`
	// 创建计算资源的时间。时间格式为ISO日期时间格式yyyy-MM-dd'T'HH:mm:ss'Z'

	CreatedTime *string `json:"created_time,omitempty"`
	// 计算资源的类型,。目前支持：sql

	ComputingResourceType *string `json:"computing_resource_type,omitempty"`
	// 与该计算资源绑定的计算单元数。设置值当前只支持16，64，256。

	CuCount *int32 `json:"cu_count,omitempty"`
	// 计算资源的收费模式。“1”表示按照CU时收费。“2”表示按照包年包月收费。

	ChargingMode *int32 `json:"charging_mode,omitempty"`
	// 计算资源类型。0：共享资源 1：专属资源

	ResourceMode *int32 `json:"resource_mode,omitempty"`
}

func (o ComputingResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputingResource struct{}"
	}

	return strings.Join([]string{"ComputingResource", string(data)}, " ")
}

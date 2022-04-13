package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateComputingResourceResponse struct {
	// 新增计算资源ID。

	ComputingResourceId *string `json:"computing_resource_id,omitempty"`
	// 新增计算资源名称。

	ComputingResourceName *string `json:"computing_resource_name,omitempty"`
	// 计算资源的类型。目前支持：sql。如果不指定，默认为sql。

	ComputingResourceType *string `json:"computing_resource_type,omitempty"`
	// 计算资源的描述信息。

	Description *string `json:"description,omitempty"`
	// 与计算资源绑定的最小计算单元个数。设置值当前只支持16，64，256。

	CuCount *int32 `json:"cu_count,omitempty"`
	// 计算资源的收费模式。只能设置为“1”，表示按照CU时收费。

	ChargingMode *int32 `json:"charging_mode,omitempty"`
	// 创建计算资源时间。时间格式为ISO日期时间格式yyyy-MM-dd'T'HH:mm:ss'Z'。

	CreatedTime    *string `json:"created_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateComputingResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComputingResourceResponse struct{}"
	}

	return strings.Join([]string{"CreateComputingResourceResponse", string(data)}, " ")
}

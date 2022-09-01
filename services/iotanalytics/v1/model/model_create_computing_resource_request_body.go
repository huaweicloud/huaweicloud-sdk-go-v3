package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateComputingResourceRequestBody struct {

	// 新建的计算资源名称，名称只能包含数字、英文字母和下划线，但不能是纯数字，且不能以下划线开头。
	ComputingResourceName string `json:"computing_resource_name" xml:"computing_resource_name"`

	// 计算资源的类型。默认为sql。
	ComputingResourceType *string `json:"computing_resource_type,omitempty" xml:"computing_resource_type"`

	// 计算资源的描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 与计算资源绑定的最小计算单元个数。设置值当前只支持16，64，256。
	CuCount int32 `json:"cu_count" xml:"cu_count"`

	// 计算资源的收费模式。只能设置为“1”，表示按照CU时收费。
	ChargingMode *int32 `json:"charging_mode,omitempty" xml:"charging_mode"`
}

func (o CreateComputingResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComputingResourceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateComputingResourceRequestBody", string(data)}, " ")
}

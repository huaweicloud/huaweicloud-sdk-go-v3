package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 解析路由结果返回结构体
type RouterDetailRespDto struct {

	// 路由ID，节点下唯一
	RouteId string `json:"route_id" xml:"route_id"`

	// 输入点所在模块的模块ID
	InputModuleId *string `json:"input_module_id,omitempty" xml:"input_module_id"`

	// 输出点所在模块的模块ID
	OutputModuleId *string `json:"output_module_id,omitempty" xml:"output_module_id"`

	// 输入点名称
	Input *string `json:"input,omitempty" xml:"input"`

	// 输出点名称
	Output *string `json:"output,omitempty" xml:"output"`

	// sql參數
	Sql *string `json:"sql,omitempty" xml:"sql"`

	// 是否可用
	Available *bool `json:"available,omitempty" xml:"available"`
}

func (o RouterDetailRespDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouterDetailRespDto struct{}"
	}

	return strings.Join([]string{"RouterDetailRespDto", string(data)}, " ")
}

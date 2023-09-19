package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandParam 扩容相关参数
type ExpandParam struct {

	// 节点组名称
	NodeGroupName string `json:"node_group_name"`

	// 扩容节点数量
	Count int32 `json:"count"`

	// 扩容时是否在新增节点上跳过执行创建集群时指定的引导操作。未填写时，默认不执行引导操作。
	SkipBootstrapScripts *bool `json:"skip_bootstrap_scripts,omitempty"`

	// 扩容后是否选择不启动扩容节点上的组件。未填写时，默认启动组件。  - true：扩容后不启动组件。 - false：扩容后启动组件。
	ScaleWithoutStart *bool `json:"scale_without_start,omitempty"`

	// 支持跨AZ场景下的扩缩容，可以传入多个AZ，逗号分隔，如果为空或“”表示非跨az场景，不进行跨AZ扩缩容校验。
	AzCode *string `json:"az_code,omitempty"`
}

func (o ExpandParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandParam struct{}"
	}

	return strings.Join([]string{"ExpandParam", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadResourceItem 资源池资源项
type WorkloadResourceItem struct {

	// 资源名称。
	ResourceName string `json:"resource_name"`

	// 资源属性值。
	ResourceValue int32 `json:"resource_value"`

	// 资源属性单位。
	ValueUnit *string `json:"value_unit,omitempty"`

	// 资源附加描述
	ResourceDescription *string `json:"resource_description,omitempty"`
}

func (o WorkloadResourceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadResourceItem struct{}"
	}

	return strings.Join([]string{"WorkloadResourceItem", string(data)}, " ")
}

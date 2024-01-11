package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OccupyResource 用户占用资源列表
type OccupyResource struct {

	// 资源项名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源属性数值
	ResourceValue *int32 `json:"resource_value,omitempty"`

	// 资源属性单位
	ValueUnit *string `json:"value_unit,omitempty"`

	// 资源附加描述
	ResourceDescription *string `json:"resource_description,omitempty"`
}

func (o OccupyResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OccupyResource struct{}"
	}

	return strings.Join([]string{"OccupyResource", string(data)}, " ")
}

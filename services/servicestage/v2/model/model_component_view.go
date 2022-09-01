package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentView struct {

	// 组件ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用ID。
	ApplicationId *string `json:"application_id,omitempty" xml:"application_id"`

	// 应用组件名称。
	Name *string `json:"name,omitempty" xml:"name"`

	Runtime *RuntimeType `json:"runtime,omitempty" xml:"runtime"`

	Category *ComponentCategory `json:"category,omitempty" xml:"category"`

	SubCategory *ComponentSubCategory `json:"sub_category,omitempty" xml:"sub_category"`

	// 组件描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 取值0或1。  0：表示正常状态。  1：表示正在删除。
	Status *int32 `json:"status,omitempty" xml:"status"`

	Source *SourceObject `json:"source,omitempty" xml:"source"`

	// 创建人。
	Creator *string `json:"creator,omitempty" xml:"creator"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 修改时间。
	UpdateTime *int64 `json:"update_time,omitempty" xml:"update_time"`
}

func (o ComponentView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentView struct{}"
	}

	return strings.Join([]string{"ComponentView", string(data)}, " ")
}

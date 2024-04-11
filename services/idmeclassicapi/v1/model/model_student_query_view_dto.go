package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StudentQueryViewDto struct {

	// 唯一编码。
	Id *string `json:"id,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 最后的修改时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// 类名。
	ClassName *string `json:"className,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 成绩。
	Grade float32 `json:"grade,omitempty"`
}

func (o StudentQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StudentQueryViewDto struct{}"
	}

	return strings.Join([]string{"StudentQueryViewDto", string(data)}, " ")
}

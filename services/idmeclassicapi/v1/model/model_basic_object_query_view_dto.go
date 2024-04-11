package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicObjectQueryViewDto struct {

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
}

func (o BasicObjectQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicObjectQueryViewDto struct{}"
	}

	return strings.Join([]string{"BasicObjectQueryViewDto", string(data)}, " ")
}

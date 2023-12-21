package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TenantQueryViewDto struct {

	// 类名。
	ClassName *string `json:"className,omitempty"`

	// 租户编码。
	Code string `json:"code"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 租户使用的数据源名称。
	DataSource string `json:"dataSource"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 失效标识。  - true：失效。  - false：未失效。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 中文名称。
	Name *string `json:"name,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantQueryViewDto `json:"tenant,omitempty"`
}

func (o TenantQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantQueryViewDto struct{}"
	}

	return strings.Join([]string{"TenantQueryViewDto", string(data)}, " ")
}

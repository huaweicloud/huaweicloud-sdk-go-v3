package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FolderQueryViewDto struct {

	// 编码。
	BusinessCode string `json:"businessCode"`

	// 类名。
	ClassName *string `json:"className,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 中文描述。
	Description *string `json:"description,omitempty"`

	// 英文描述。
	DescriptionEn *string `json:"descriptionEn,omitempty"`

	// 失效标识。  - true：失效。  - false：未失效。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	// 扩展属性映射集。
	ExtAttrMap *interface{} `json:"extAttrMap,omitempty"`

	// 扩展属性列表。
	ExtAttrs *[]ExaValueViewDto `json:"extAttrs,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 中文名称。
	Name string `json:"name"`

	// 英文名称。
	NameEn *string `json:"nameEn,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantQueryViewDto `json:"tenant,omitempty"`

	// 类别。
	Type *string `json:"type,omitempty"`
}

func (o FolderQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FolderQueryViewDto struct{}"
	}

	return strings.Join([]string{"FolderQueryViewDto", string(data)}, " ")
}

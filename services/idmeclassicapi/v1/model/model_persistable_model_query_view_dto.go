package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelQueryViewDto struct {

	// 类名。
	ClassName *string `json:"className,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 失效标识。  - true：失效。  - false：未失效。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	// 扩展属性映射集。
	ExtAttrMap *interface{} `json:"extAttrMap,omitempty"`

	// 扩展属性列表。
	ExtAttrs *[]ExaValueViewDto `json:"extAttrs,omitempty"`

	Folder *FolderQueryViewDto `json:"folder,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantQueryViewDto `json:"tenant,omitempty"`
}

func (o PersistableModelQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelQueryViewDto struct{}"
	}

	return strings.Join([]string{"PersistableModelQueryViewDto", string(data)}, " ")
}

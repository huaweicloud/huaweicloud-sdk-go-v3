package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelUpdateDto struct {

	// 访问控制列表。
	AclEntry *string `json:"aclEntry,omitempty"`

	// 分类属性。
	ClsAttrs *[]interface{} `json:"clsAttrs,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 失效标识。  - true：失效。  - false：未失效。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	// 扩展属性映射集。
	ExtAttrMap *interface{} `json:"extAttrMap,omitempty"`

	// 扩展属性列表。
	ExtAttrs *[]ExaValueParamDto `json:"extAttrs,omitempty"`

	Folder *ObjectReferenceParamDto `json:"folder,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	LifecycleState *ObjectReferenceParamDto `json:"lifecycleState,omitempty"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// 设置NULL值的属性名称。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// 拥有者。
	Owner *string `json:"owner,omitempty"`

	ParentNode *ObjectReferenceParamDto `json:"parentNode,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 示例模型的唯一键属性。
	UniqueKey *string `json:"uniqueKey,omitempty"`
}

func (o PersistableModelUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelUpdateDto struct{}"
	}

	return strings.Join([]string{"PersistableModelUpdateDto", string(data)}, " ")
}

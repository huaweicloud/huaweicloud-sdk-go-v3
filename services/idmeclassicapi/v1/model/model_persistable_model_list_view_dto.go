package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelListViewDto struct {

	// 访问控制列表。
	AclEntry *string `json:"aclEntry,omitempty"`

	// 类名。
	ClassName *string `json:"className,omitempty"`

	// 分类属性。
	ClsAttrs *[]interface{} `json:"clsAttrs,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 失效标识。  - true：失效。  - false：未失效。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	Folder *ObjectReferenceViewDto `json:"folder,omitempty"`

	// 用于存储当前节点全路径。
	FullPath *string `json:"fullPath,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 是否为叶子节点。 - true：是叶子节点。 - false：不是叶子节点。
	LeafFlag *bool `json:"leafFlag,omitempty"`

	LifecycleState *ObjectReferenceViewDto `json:"lifecycleState,omitempty"`

	LifecycleTemplate *ObjectReferenceViewDto `json:"lifecycleTemplate,omitempty"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// 拥有者。
	Owner *string `json:"owner,omitempty"`

	ParentNode *ObjectReferenceViewDto `json:"parentNode,omitempty"`

	// 用于存储当前节点原始全路径。
	RawFullPath *string `json:"rawFullPath,omitempty"`

	// 软删除标识，参数值为0或1。 - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 系统版本。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	RootNode *ObjectReferenceViewDto `json:"rootNode,omitempty"`

	Tenant *ObjectReferenceViewDto `json:"tenant,omitempty"`

	// 示例模型的唯一键属性。
	UniqueKey *string `json:"uniqueKey,omitempty"`
}

func (o PersistableModelListViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelListViewDto struct{}"
	}

	return strings.Join([]string{"PersistableModelListViewDto", string(data)}, " ")
}

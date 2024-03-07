package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelSaveDto struct {

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// 设置NULL值的属性名称。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *ObjectReferenceParamDto `json:"tenant,omitempty"`

	// 示例模型的唯一键属性。
	UniqueKey *string `json:"uniqueKey,omitempty"`
}

func (o PersistableModelSaveDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelSaveDto struct{}"
	}

	return strings.Join([]string{"PersistableModelSaveDto", string(data)}, " ")
}

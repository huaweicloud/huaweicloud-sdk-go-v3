package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelSaveAsDto struct {

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 修改者。
	Modifier *string `json:"modifier,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *ObjectReferenceParamDto `json:"tenant,omitempty"`

	// 源模型编号。
	SourceEntityNumber *string `json:"sourceEntityNumber,omitempty"`

	// 源实例的唯一标识(单实例为ID，版本实例为versionId)。
	SourceInstanceId string `json:"sourceInstanceId"`

	// 置空字段数组。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// 要保存的属性。
	EntityToSave *interface{} `json:"entityToSave,omitempty"`

	// 要保存的结果。
	EntityToReturn *interface{} `json:"entityToReturn,omitempty"`

	// 唯一键。
	UniqueKey *string `json:"uniqueKey,omitempty"`
}

func (o PersistableModelSaveAsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelSaveAsDto struct{}"
	}

	return strings.Join([]string{"PersistableModelSaveAsDto", string(data)}, " ")
}

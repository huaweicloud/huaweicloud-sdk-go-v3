package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelSaveAsDto struct {
	Branch *VersionModelBranchSaveAsDto `json:"branch,omitempty"`

	// 检出时间。
	CheckOutTime *string `json:"checkOutTime,omitempty"`

	// 检出人。
	CheckOutUserName *string `json:"checkOutUserName,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 需要返回的实体。
	EntityToReturn *interface{} `json:"entityToReturn,omitempty"`

	// 需要保存的实体。
	EntityToSave *interface{} `json:"entityToSave,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 关键信息资产ID。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	Master *VersionModelMasterSaveAsDto `json:"master,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 中文名称。
	Name *string `json:"name,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 安全密级。 - INTERNAL：内部公开。 - SECRET：秘密。 - CONFIDENTIAL：机密。 - TOP_SECRET：绝密。
	SecurityLevel *string `json:"securityLevel,omitempty"`

	// version.唯一编码。
	SourceInstanceId string `json:"sourceInstanceId"`

	// master.唯一编码。
	SourceMasterId *string `json:"sourceMasterId,omitempty"`

	Tenant *ObjectReferenceParamDto `json:"tenant,omitempty"`

	// 是否已检出。 - true：已检出。 - false：未检出。
	WorkingCopy *bool `json:"workingCopy,omitempty"`
}

func (o VersionModelSaveAsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelSaveAsDto struct{}"
	}

	return strings.Join([]string{"VersionModelSaveAsDto", string(data)}, " ")
}

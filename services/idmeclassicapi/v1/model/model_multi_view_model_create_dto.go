package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiViewModelCreateDto struct {
	Item *ObjectReferenceParamDto `json:"item,omitempty"`

	Branch *VersionModelBranchCreateDto `json:"branch,omitempty"`

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

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 关键信息资产ID。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	Master *VersionModelMasterCreateDto `json:"master,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 中文名称。
	Name *string `json:"name,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 安全密级。 - INTERNAL：内部公开。 - SECRET：秘密。 - CONFIDENTIAL：机密。 - TOP_SECRET：绝密。
	SecurityLevel *string `json:"securityLevel,omitempty"`

	Tenant *ObjectReferenceParamDto `json:"tenant,omitempty"`

	// 是否已检出。 - true：已检出。 - false：未检出。
	WorkingCopy *bool `json:"workingCopy,omitempty"`
}

func (o MultiViewModelCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiViewModelCreateDto struct{}"
	}

	return strings.Join([]string{"MultiViewModelCreateDto", string(data)}, " ")
}

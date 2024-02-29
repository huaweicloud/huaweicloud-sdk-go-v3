package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModel struct {
	Branch *VersionModelBranch `json:"branch,omitempty"`

	// 检出时间。
	CheckOutTime *string `json:"checkOutTime,omitempty"`

	// 检出用户名称。
	CheckOutUserName *string `json:"checkOutUserName,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 唯一标识。
	Id *int64 `json:"id,omitempty"`

	// 关键信息资产ID。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	Master *VersionModelMaster `json:"master,omitempty"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// 中文名称。
	Name *string `json:"name,omitempty"`

	// 设置NULL值的属性名称。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 安全密级。 - INTERNAL：内部公开。 - SECRET：秘密。 - CONFIDENTIAL：机密。 - TOP_SECRET：绝密。
	SecurityLevel *string `json:"securityLevel,omitempty"`
}

func (o VersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModel struct{}"
	}

	return strings.Join([]string{"VersionModel", string(data)}, " ")
}

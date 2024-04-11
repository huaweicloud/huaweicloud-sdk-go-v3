package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StudentViewDto struct {

	// 类名。
	ClassName *string `json:"className,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 成绩。
	Grade float32 `json:"grade,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 关键信息资产ID。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// 最新更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 软删除标识，参数值为0或1。  - 0：表示未删除。  - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 系统版本。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// 安全密级。  - INTERNAL：内部公开。  - SECRET：秘密。  - CONFIDENTIAL：机密。  - TOP_SECRET：绝密。
	SecurityLevel *string `json:"securityLevel,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`
}

func (o StudentViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StudentViewDto struct{}"
	}

	return strings.Join([]string{"StudentViewDto", string(data)}, " ")
}

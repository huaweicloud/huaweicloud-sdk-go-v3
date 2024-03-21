package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TenantHistoryViewDto struct {

	// 类名。
	ClassName *string `json:"className,omitempty"`

	// 租户编码。
	Code string `json:"code"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 租户使用的数据源名称。
	DataSource string `json:"dataSource"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 失效标识。  - true：失效。  - false：未失效。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// KIA密级。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// 中文名称。
	Name *string `json:"name,omitempty"`

	// 软删除标识，参数值为0或1。 - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 操作类型，用于存储MONGO。 - CASCADE：级联。 - CREATE：级联。 - DELETE：创建。 - LOGICALDELETE：软删除。 - UPDATE：更新。
	RdmOperationType *string `json:"rdmOperationType,omitempty"`

	// 系统版本。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// 安全密级。 - INTERNAL：内部公开。 - SECRET：秘密。 - CONFIDENTIAL：机密。 - TOP_SECRET：绝密。
	SecurityLevel *string `json:"securityLevel,omitempty"`

	Tenant *TenantHistoryViewDto `json:"tenant,omitempty"`
}

func (o TenantHistoryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantHistoryViewDto struct{}"
	}

	return strings.Join([]string{"TenantHistoryViewDto", string(data)}, " ")
}

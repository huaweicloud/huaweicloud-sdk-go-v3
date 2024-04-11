package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryDataModelHistoryViewDto struct {

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 创建人。
	Creator *string `json:"creator,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 修改时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 系统版本。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// 操作类型。 - CREATE：创建操作。 - UPDATE：更新操作。 - LOGICALDELETE：软删除操作。 - DELETE：删除操作。 - CASCADE：级联操作。
	RdmOperationType *string `json:"rdmOperationType,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 删除标志。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	Tenant *TenantHistoryViewDto `json:"tenant,omitempty"`

	// 类名称。
	ClassName *string `json:"className,omitempty"`
}

func (o HistoryDataModelHistoryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryDataModelHistoryViewDto struct{}"
	}

	return strings.Join([]string{"HistoryDataModelHistoryViewDto", string(data)}, " ")
}

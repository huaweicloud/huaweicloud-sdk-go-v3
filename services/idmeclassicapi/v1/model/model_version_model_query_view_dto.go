package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelQueryViewDto struct {
	Branch *VersionObjectBranchQueryViewDto `json:"branch,omitempty"`

	// 检出时间。
	CheckOutTime *string `json:"checkOutTime,omitempty"`

	// 检出人。
	CheckOutUserName *string `json:"checkOutUserName,omitempty"`

	// 类名。
	ClassName *string `json:"className,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 迭代版本。
	Iteration *int32 `json:"iteration,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 是否为最新版本。 - true：是最新版本。 - false：不是最新版本。
	Latest *bool `json:"latest,omitempty"`

	// 是否为最新迭代版本。 - true：是最新迭代版本。 - false：不是最新迭代版本。
	LatestIteration *bool `json:"latestIteration,omitempty"`

	// 是否为最新修订版本。 - true：是最新修订版本。 - false：不是最新修订版本。
	LatestVersion *bool `json:"latestVersion,omitempty"`

	Master *MasterObjectQueryViewDto `json:"master,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 中文名称。
	Name *string `json:"name,omitempty"`

	// 前序版本实例ID。
	PreVersionId *string `json:"preVersionId,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantQueryViewDto `json:"tenant,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`

	// 业务版本内码。
	VersionCode *int32 `json:"versionCode,omitempty"`

	// 是否已检出。 - true：已检出。 - false：未检出。
	WorkingCopy *bool `json:"workingCopy,omitempty"`

	WorkingState *WorkingState `json:"workingState,omitempty"`
}

func (o VersionModelQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelQueryViewDto struct{}"
	}

	return strings.Join([]string{"VersionModelQueryViewDto", string(data)}, " ")
}

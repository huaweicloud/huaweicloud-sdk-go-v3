package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiViewModelBranchViewDto struct {

	// 类名。
	ClassName *string `json:"className,omitempty"`

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 最后更新时间。
	LastUpdateTime *interface{} `json:"lastUpdateTime,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 软删除标识，参数值为0或1。 - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// 系统版本。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`
}

func (o MultiViewModelBranchViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiViewModelBranchViewDto struct{}"
	}

	return strings.Join([]string{"MultiViewModelBranchViewDto", string(data)}, " ")
}

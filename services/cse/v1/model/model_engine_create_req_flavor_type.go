package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineCreateReqFlavorType 网关规格类型
type EngineCreateReqFlavorType struct {

	// 网关节点规格
	NodeFlavor *[]string `json:"nodeFlavor,omitempty"`

	// 网关规格
	Flavor *string `json:"flavor,omitempty"`

	// 可用区前缀
	AvailablePrefix *string `json:"availablePrefix,omitempty"`

	// 可用区CPU内存
	AvailableCpuMemory *string `json:"availableCpuMemory,omitempty"`

	// 引擎类型
	SpecType *string `json:"specType,omitempty"`

	// 是否为线性
	Linear *bool `json:"linear,omitempty"`

	// 网关证书规模
	LicenseAmount *int32 `json:"licenseAmount,omitempty"`

	// 网关节点数限制
	NodeLimit *string `json:"nodeLimit,omitempty"`

	// 网关规格id
	Id *string `json:"id,omitempty"`

	// 网关规格
	MicroGatewayFlavor *string `json:"microGatewayFlavor,omitempty"`

	// 网关是否禁用
	Disable *bool `json:"disable,omitempty"`

	// 网关节点类型
	Spec *string `json:"spec,omitempty"`

	// 云服务类型
	CloudServiceType *string `json:"cloudServiceType,omitempty"`

	// 当前规格
	Currentflavor *string `json:"currentflavor,omitempty"`
}

func (o EngineCreateReqFlavorType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineCreateReqFlavorType struct{}"
	}

	return strings.Join([]string{"EngineCreateReqFlavorType", string(data)}, " ")
}

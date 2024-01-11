package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsExtraSpecs
type OsExtraSpecs struct {

	// 此参数是Region级配置，某个AZ没有在cond:operation:az参数中配置时默认使用此参数的取值。不配置或无此参数时等同于“normal”。
	Condoperationstatus *string `json:"cond:operation:status,omitempty"`

	// 边缘实例类型的代数。
	Ecsgeneration *string `json:"ecs:generation,omitempty"`

	// 边缘实例规格的分类。
	Ecsperformancetype *string `json:"ecs:performancetype,omitempty"`

	// 虚拟化类型。
	EcsvirtualizationEnvTypes *string `json:"ecs:virtualization_env_types,omitempty"`

	// 此参数是规格的CPU详细信息。
	Infocpuname *string `json:"info:cpu:name,omitempty"`

	// 此参数是规格的GPU详细信息。
	Infogpuname *string `json:"info:gpu:name,omitempty"`

	// P1型本地直通GPU的型号和数量，参数值可设置为“nvidia-p100:1”，表示使用该规格创建的边缘实例将占用1张NVIDIA P100显卡。
	PciPassthroughalias *string `json:"pci_passthrough:alias,omitempty"`

	// 显卡是否直通。 值为“true”，表示GPU直通。
	PciPassthroughenableGpu *string `json:"pci_passthrough:enable_gpu,omitempty"`

	// G1型和G2型边缘实例应用的技术，包括GPU虚拟化和GPU直通。
	PciPassthroughgpuSpecs *string `json:"pci_passthrough:gpu_specs,omitempty"`

	// 资源类型，resource_type是为了区分边缘实例的物理主机类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 网络约束  支持网络特性，不配置时以控制台默认配置为准。
	Condnetwork *string `json:"cond:network,omitempty"`
}

func (o OsExtraSpecs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsExtraSpecs struct{}"
	}

	return strings.Join([]string{"OsExtraSpecs", string(data)}, " ")
}

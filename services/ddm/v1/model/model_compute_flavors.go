package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputeFlavors struct {
	// 规格id。

	Id *string `json:"id,omitempty"`
	// 资源类型编码。

	TypeCode *string `json:"typeCode,omitempty"`
	// DDM内部记录的虚机规格类型。

	Code *string `json:"code,omitempty"`
	// iaas记录的虚机规格类型。

	IaasCode *string `json:"iaasCode,omitempty"`
	// cpu核数。

	Cpu *string `json:"cpu,omitempty"`
	// 内存大小,单位:G。

	Mem *string `json:"mem,omitempty"`
	// 最大连接数。

	MaxConnections *string `json:"maxConnections,omitempty"`
	// 计算资源服务类型。

	ServerType *string `json:"serverType,omitempty"`
	// 计算资源架构类型，目前分X86和ARM两种。

	Architecture *string `json:"architecture,omitempty"`
	// 可用区状态。

	AzStatus *interface{} `json:"azStatus,omitempty"`
	// 局点状态。

	RegionStatus *string `json:"regionStatus,omitempty"`
	// 计算资源架构类型，目前分X86和ARM两种。

	GroupType *string `json:"groupType,omitempty"`
	// 服务引擎类型。

	DbType *string `json:"dbType,omitempty"`
	// 扩展字段，目前存储可用区相关信息。

	ExtendFields *interface{} `json:"extendFields,omitempty"`
}

func (o ComputeFlavors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeFlavors struct{}"
	}

	return strings.Join([]string{"ComputeFlavors", string(data)}, " ")
}

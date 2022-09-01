package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputeFlavors struct {

	// 规格id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 资源类型编码。
	TypeCode *string `json:"typeCode,omitempty" xml:"typeCode"`

	// DDM内部记录的虚机规格类型。
	Code *string `json:"code,omitempty" xml:"code"`

	// iaas记录的虚机规格类型。
	IaasCode *string `json:"iaasCode,omitempty" xml:"iaasCode"`

	// cpu核数。
	Cpu *string `json:"cpu,omitempty" xml:"cpu"`

	// 内存大小,单位:G。
	Mem *string `json:"mem,omitempty" xml:"mem"`

	// 最大连接数。
	MaxConnections *string `json:"maxConnections,omitempty" xml:"maxConnections"`

	// 计算资源服务类型。
	ServerType *string `json:"serverType,omitempty" xml:"serverType"`

	// 计算资源架构类型，目前分X86和ARM两种。
	Architecture *string `json:"architecture,omitempty" xml:"architecture"`

	// 可用区状态。
	AzStatus *interface{} `json:"azStatus,omitempty" xml:"azStatus"`

	// 局点状态。
	RegionStatus *string `json:"regionStatus,omitempty" xml:"regionStatus"`

	// 计算资源架构类型，目前分X86和ARM两种。
	GroupType *string `json:"groupType,omitempty" xml:"groupType"`

	// 服务引擎类型。
	DbType *string `json:"dbType,omitempty" xml:"dbType"`

	// 扩展字段，目前存储可用区相关信息。
	ExtendFields *interface{} `json:"extendFields,omitempty" xml:"extendFields"`
}

func (o ComputeFlavors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeFlavors struct{}"
	}

	return strings.Join([]string{"ComputeFlavors", string(data)}, " ")
}

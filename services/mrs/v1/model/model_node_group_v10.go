package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NodeGroupV10 struct {

	// 节点组名。
	GroupName *string `json:"GroupName,omitempty"`

	// 节点数量，取值范围0～500，Master节点和Core节点数量至少为1，Core与Task节点总数最大为500个。
	NodeNum *int32 `json:"NodeNum,omitempty"`

	// 节点的实例规格。
	NodeSize *string `json:"NodeSize,omitempty"`

	// 节点实例规格ID。
	NodeSpecId *string `json:"NodeSpecId,omitempty"`

	// 节点虚拟机产品ID。
	VmProductId *string `json:"VmProductId,omitempty"`

	// 节点虚拟机产品规格。
	VmSpecCode *string `json:"VmSpecCode,omitempty"`

	// 节点实例产品ID。
	NodeProductId *string `json:"NodeProductId,omitempty"`

	// 节点系统盘大小，不可配置，默认为40GB。
	RootVolumeSize *int32 `json:"RootVolumeSize,omitempty"`

	// 节点系统盘的产品ID。
	RootVolumeProductId *string `json:"RootVolumeProductId,omitempty"`

	// 节点系统盘的类型。
	RootVolumeType *string `json:"RootVolumeType,omitempty"`

	// 节点系统盘产品规格。
	RootVolumeResourceSpecCode *string `json:"RootVolumeResourceSpecCode,omitempty"`

	// 节点系统盘产品类型。
	RootVolumeResourceType *string `json:"RootVolumeResourceType,omitempty"`

	// 节点数据磁盘存储类别，目前支持SATA、SAS和SSD。  - SATA：普通IO - SAS：高IO - SSD：超高IO
	DataVolumeType *NodeGroupV10DataVolumeType `json:"DataVolumeType,omitempty"`

	// 节点数据磁盘存储数目。
	DataVolumeCount *int32 `json:"DataVolumeCount,omitempty"`

	// 节点数据磁盘存储大小。
	DataVolumeSize *int32 `json:"DataVolumeSize,omitempty"`

	// 节点数据磁盘的产品ID。
	DataVolumeProductId *string `json:"DataVolumeProductId,omitempty"`

	// 节点数据磁盘的产品规格。
	DataVolumeResourceSpecCode *string `json:"DataVolumeResourceSpecCode,omitempty"`

	// 节点数据磁盘的产品类型。
	DataVolumeResourceType *string `json:"DataVolumeResourceType,omitempty"`
}

func (o NodeGroupV10) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeGroupV10 struct{}"
	}

	return strings.Join([]string{"NodeGroupV10", string(data)}, " ")
}

type NodeGroupV10DataVolumeType struct {
	value string
}

type NodeGroupV10DataVolumeTypeEnum struct {
	SATA  NodeGroupV10DataVolumeType
	SAS   NodeGroupV10DataVolumeType
	SSD   NodeGroupV10DataVolumeType
	GPSSD NodeGroupV10DataVolumeType
}

func GetNodeGroupV10DataVolumeTypeEnum() NodeGroupV10DataVolumeTypeEnum {
	return NodeGroupV10DataVolumeTypeEnum{
		SATA: NodeGroupV10DataVolumeType{
			value: "SATA",
		},
		SAS: NodeGroupV10DataVolumeType{
			value: "SAS",
		},
		SSD: NodeGroupV10DataVolumeType{
			value: "SSD",
		},
		GPSSD: NodeGroupV10DataVolumeType{
			value: "GPSSD",
		},
	}
}

func (c NodeGroupV10DataVolumeType) Value() string {
	return c.value
}

func (c NodeGroupV10DataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeGroupV10DataVolumeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetClusterFlavorSpecsRequest Request Object
type GetClusterFlavorSpecsRequest struct {

	// **参数解释**： 该参数用于按集群架构查询可售卖规格 **取值范围**： - VirtualMachine: CCE集群 - ARM64: 鲲鹏集群
	ClusterType GetClusterFlavorSpecsRequestClusterType `json:"clusterType"`
}

func (o GetClusterFlavorSpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetClusterFlavorSpecsRequest struct{}"
	}

	return strings.Join([]string{"GetClusterFlavorSpecsRequest", string(data)}, " ")
}

type GetClusterFlavorSpecsRequestClusterType struct {
	value string
}

type GetClusterFlavorSpecsRequestClusterTypeEnum struct {
	VIRTUAL_MACHINE GetClusterFlavorSpecsRequestClusterType
	ARM64           GetClusterFlavorSpecsRequestClusterType
}

func GetGetClusterFlavorSpecsRequestClusterTypeEnum() GetClusterFlavorSpecsRequestClusterTypeEnum {
	return GetClusterFlavorSpecsRequestClusterTypeEnum{
		VIRTUAL_MACHINE: GetClusterFlavorSpecsRequestClusterType{
			value: "VirtualMachine",
		},
		ARM64: GetClusterFlavorSpecsRequestClusterType{
			value: "ARM64",
		},
	}
}

func (c GetClusterFlavorSpecsRequestClusterType) Value() string {
	return c.value
}

func (c GetClusterFlavorSpecsRequestClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetClusterFlavorSpecsRequestClusterType) UnmarshalJSON(b []byte) error {
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

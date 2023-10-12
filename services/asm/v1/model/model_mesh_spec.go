package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MeshSpec 网格参数定义。
type MeshSpec struct {

	// 网格控制面组件所在的region。 创建企业版网格时需要填写该参数，选择控制面组件所在的region； 创建基础版时网格组件安装在用户所提供的集群中，不需要填写该参数。
	Region *string `json:"region,omitempty"`

	// 网格类型：  Managed：企业版网格  InCluster：基础版网格
	Type MeshSpecType `json:"type"`

	// 网格版本
	Version string `json:"version"`

	ExtendParams *MeshExtendParams `json:"extendParams,omitempty"`
}

func (o MeshSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeshSpec struct{}"
	}

	return strings.Join([]string{"MeshSpec", string(data)}, " ")
}

type MeshSpecType struct {
	value string
}

type MeshSpecTypeEnum struct {
	MANAGED    MeshSpecType
	IN_CLUSTER MeshSpecType
}

func GetMeshSpecTypeEnum() MeshSpecTypeEnum {
	return MeshSpecTypeEnum{
		MANAGED: MeshSpecType{
			value: "Managed",
		},
		IN_CLUSTER: MeshSpecType{
			value: "InCluster",
		},
	}
}

func (c MeshSpecType) Value() string {
	return c.value
}

func (c MeshSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MeshSpecType) UnmarshalJSON(b []byte) error {
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

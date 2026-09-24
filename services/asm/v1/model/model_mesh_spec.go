package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MeshSpec 网格参数定义
type MeshSpec struct {

	// 网格类型。 取值范围： - InCluster: 集群内控制平面形态，基础版网格取值为InCluster。目前仅支持该类型。
	Type MeshSpecType `json:"type"`

	// **参数解释：** 网格版本，与Istio社区基线版本保持一致，建议选择最新商用版本。  在ASM控制台支持创建多种版本的网格。可登录ASM控制台创建网格，在“版本”处获取到网格版本。 **约束限制：** 格式必须为：vX.Y[.Z[-rN]]，例如 v1.18，v1.18.7，v1.18.7-r1 都将创建1.18版本的网格 - X: 对应社区Istio的主要版本 - Y: 对应社区Istio的次要版本 - Z: 对应社区Istio的补丁版本 - N: 对应ASM补丁版本  **取值范围：** 不涉及 **默认取值：** - 若不配置，默认创建最新版本的网格。 - 若指定网格基线版本但是不指定具体r版本，则系统默认选择对应网格版本的最新r版本。建议不指定具体r版本由系统选择最新版本。
	Version *string `json:"version,omitempty"`

	ExtendParams *MeshExtendParams `json:"extendParams"`

	// 网格是否支持IPV6
	Ipv6Enable *bool `json:"ipv6Enable,omitempty"`

	// 网格资源标签。如果需要配置资源标签，请确认当前region的TMS服务已上线。
	Tags *[]MeshTags `json:"tags,omitempty"`

	Config *MeshConfig `json:"config,omitempty"`
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
	IN_CLUSTER MeshSpecType
}

func GetMeshSpecTypeEnum() MeshSpecTypeEnum {
	return MeshSpecTypeEnum{
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

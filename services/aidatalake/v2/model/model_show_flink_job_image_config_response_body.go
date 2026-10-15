package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFlinkJobImageConfigResponseBody **参数解释**：Flink镜像配置，包含镜像ID和版本信息。
type ShowFlinkJobImageConfigResponseBody struct {

	// **参数解释**：镜像ID，Flink作业运行所使用的镜像。 **取值范围**：长度为2~64个字符。
	ImageId string `json:"image_id"`

	// **参数解释**：镜像版本ID，Flink作业运行所使用的镜像版本。 **取值范围**：长度为2~64个字符。
	ImageVersionId string `json:"image_version_id"`

	// **参数解释**：Flink作业的版本信息，格式为AIDataLake.<平台版本>(Flink<Flink版本>,JRE<JRE版本>)，例如AIDataLake.26.9.0(Flink1.20.0,JRE21)。 **取值范围**：长度为1~512个字符。
	ReleaseVersion *string `json:"release_version,omitempty"`

	// **参数解释**：Flink作业使用的镜像版本类型。 **取值范围**：   - STABLE：稳定版。   - RECOMMEND：推荐版。   - BETA：测试版。
	ImageVersionType *ShowFlinkJobImageConfigResponseBodyImageVersionType `json:"image_version_type,omitempty"`
}

func (o ShowFlinkJobImageConfigResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobImageConfigResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobImageConfigResponseBody", string(data)}, " ")
}

type ShowFlinkJobImageConfigResponseBodyImageVersionType struct {
	value string
}

type ShowFlinkJobImageConfigResponseBodyImageVersionTypeEnum struct {
	STABLE    ShowFlinkJobImageConfigResponseBodyImageVersionType
	RECOMMEND ShowFlinkJobImageConfigResponseBodyImageVersionType
	BETA      ShowFlinkJobImageConfigResponseBodyImageVersionType
}

func GetShowFlinkJobImageConfigResponseBodyImageVersionTypeEnum() ShowFlinkJobImageConfigResponseBodyImageVersionTypeEnum {
	return ShowFlinkJobImageConfigResponseBodyImageVersionTypeEnum{
		STABLE: ShowFlinkJobImageConfigResponseBodyImageVersionType{
			value: "STABLE",
		},
		RECOMMEND: ShowFlinkJobImageConfigResponseBodyImageVersionType{
			value: "RECOMMEND",
		},
		BETA: ShowFlinkJobImageConfigResponseBodyImageVersionType{
			value: "BETA",
		},
	}
}

func (c ShowFlinkJobImageConfigResponseBodyImageVersionType) Value() string {
	return c.value
}

func (c ShowFlinkJobImageConfigResponseBodyImageVersionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlinkJobImageConfigResponseBodyImageVersionType) UnmarshalJSON(b []byte) error {
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

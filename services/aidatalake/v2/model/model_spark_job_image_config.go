package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJobImageConfig **参数解释**：Spark镜像配置，用于指定Spark作业运行所需的镜像信息。 **约束限制**：不涉及。
type SparkJobImageConfig struct {

	// **参数解释**：镜像ID，用于指定Spark作业运行所需的镜像。 **约束限制**：不涉及。 **取值范围**：长度为2~64个字符。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`

	// **参数解释**：镜像版本，用于指定Spark作业运行所需的镜像版本。 **约束限制**：不涉及。 **取值范围**：长度为2~64个字符。 **默认取值**：不涉及。
	ImageVersionId string `json:"image_version_id"`
}

func (o SparkJobImageConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobImageConfig struct{}"
	}

	return strings.Join([]string{"SparkJobImageConfig", string(data)}, " ")
}

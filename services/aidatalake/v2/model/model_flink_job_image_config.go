package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobImageConfig **参数解释**：Flink镜像配置，用于指定Flink作业运行所需的镜像信息。 **约束限制**：不涉及。
type FlinkJobImageConfig struct {

	// **参数解释**：镜像ID，用于指定Flink作业运行所需的镜像。须为平台镜像管理中已存在的镜像，传入不存在的镜像时接口返回错误，可通过平台镜像管理功能获取。 **约束限制**：须与image_version_id配合使用，二者对应的镜像版本须已存在。 **取值范围**：长度为2~64个字符，字符组成不限制。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`

	// **参数解释**：镜像版本ID，用于指定Flink作业运行所需的镜像版本。须为平台镜像管理中已存在的镜像版本，传入不存在的镜像版本时接口返回错误，可通过平台镜像管理功能获取。 **约束限制**：须与image_id配合使用，二者对应的镜像版本须已存在。 **取值范围**：长度为2~64个字符，字符组成不限制。 **默认取值**：不涉及。
	ImageVersionId string `json:"image_version_id"`
}

func (o FlinkJobImageConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobImageConfig struct{}"
	}

	return strings.Join([]string{"FlinkJobImageConfig", string(data)}, " ")
}

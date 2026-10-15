package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayJobConfig **参数解释**：Ray作业类型的配置信息。key为配置项名称，类型为String，value为配置项值，类型为Object。 **约束限制**：不涉及。
type RayJobConfig struct {

	// **参数解释**：运行作业主脚本。 **约束限制**：不涉及。 **取值范围**：长度为1~102400个字符。 **默认取值**：不涉及。
	Entrypoint string `json:"entrypoint"`

	JobConfig *InnerRayJobConfig `json:"job_config,omitempty"`

	RuntimeEnv *RuntimeEnv `json:"runtime_env,omitempty"`

	Image *V2ImageBriefInfo `json:"image,omitempty"`

	ResourceConfig *RayResourceConfig `json:"resource_config,omitempty"`

	RayjobStrategy *RayJobStrategy `json:"rayjob_strategy,omitempty"`

	// **参数解释**：用户自定义作业元数据（键值对）。 **约束限制**：键值对数量为0~255个。key和value的值不能为空，长度不超过255个字符 。
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (o RayJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayJobConfig struct{}"
	}

	return strings.Join([]string{"RayJobConfig", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataframeJobConfig **参数解释**：Aura dataframe Job作业类型的配置信息 。  **约束限制**：不涉及。 **取值范围**：不涉及 **默认取值**：不涉及。
type DataframeJobConfig struct {

	// **参数解释**：运行作业主脚本。  **约束限制**：不涉及。 **取值范围**：长度为1~1024个字符。 **默认取值**：不涉及。
	Entrypoint string `json:"entrypoint"`

	JobConfig *ResourceConfig `json:"job_config,omitempty"`

	RuntimeEnv *JobRuntimeEnv `json:"runtime_env,omitempty"`
}

func (o DataframeJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataframeJobConfig struct{}"
	}

	return strings.Join([]string{"DataframeJobConfig", string(data)}, " ")
}

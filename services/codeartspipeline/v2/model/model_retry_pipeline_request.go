package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryPipelineRequest 重试运行流水线请求体。
type RetryPipelineRequest struct {

	// **参数解释**： 仓库HTTPS地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RepoHttpsUrl *string `json:"repo_https_url,omitempty"`

	// **参数解释**： 流水线任务运行ID列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	JobRunIds *[]string `json:"job_run_ids,omitempty"`
}

func (o RetryPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryPipelineRequest struct{}"
	}

	return strings.Join([]string{"RetryPipelineRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobItem 批量删除训练作业时的作业标识信息。
type BatchDeleteJobItem struct {

	// **参数解释**：训练作业类型。 **取值范围**： - job - edge_job - mrs_job - hetero_job - autosearch_job - diag_job - visualization_job - federated_pool_job
	Kind string `json:"kind"`

	Metadata *BatchDeleteJobMetadata `json:"metadata"`
}

func (o BatchDeleteJobItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobItem struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobItem", string(data)}, " ")
}

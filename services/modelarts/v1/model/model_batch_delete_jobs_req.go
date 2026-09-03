package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobsReq 批量删除训练作业请求体。
type BatchDeleteJobsReq struct {

	// **参数解释**：待删除的训练作业列表。 **约束限制**：列表元素数量不超过100，且所有作业必须属于同一工作空间。 **取值范围**：不涉及。
	Jobs []BatchDeleteJobItem `json:"jobs"`
}

func (o BatchDeleteJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsReq", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteClusterCnResponse Response Object
type BatchDeleteClusterCnResponse struct {

	// **参数解释**： 批量删除CN节点任务ID信息。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteClusterCnResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterCnResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterCnResponse", string(data)}, " ")
}

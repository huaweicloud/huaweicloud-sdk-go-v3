package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkClusterResponse Response Object
type ShrinkClusterResponse struct {

	// **参数解释**： 缩容的任务ID信息。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkClusterResponse struct{}"
	}

	return strings.Join([]string{"ShrinkClusterResponse", string(data)}, " ")
}

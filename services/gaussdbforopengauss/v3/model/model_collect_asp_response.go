package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectAspResponse Response Object
type CollectAspResponse struct {

	// **参数解释**: 任务ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectAspResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectAspResponse struct{}"
	}

	return strings.Join([]string{"CollectAspResponse", string(data)}, " ")
}

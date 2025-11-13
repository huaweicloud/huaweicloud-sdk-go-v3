package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePublicZonesTaskResponse Response Object
type BatchCreatePublicZonesTaskResponse struct {

	// **参数解释：** 批量操作任务的ID。 **取值范围：** 不涉及。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreatePublicZonesTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicZonesTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicZonesTaskResponse", string(data)}, " ")
}

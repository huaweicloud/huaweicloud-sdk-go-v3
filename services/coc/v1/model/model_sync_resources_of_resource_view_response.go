package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncResourcesOfResourceViewResponse Response Object
type SyncResourcesOfResourceViewResponse struct {

	// **参数解释：** 任务id。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SyncResourcesOfResourceViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncResourcesOfResourceViewResponse struct{}"
	}

	return strings.Join([]string{"SyncResourcesOfResourceViewResponse", string(data)}, " ")
}

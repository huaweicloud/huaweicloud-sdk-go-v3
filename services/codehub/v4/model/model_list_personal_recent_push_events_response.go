package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPersonalRecentPushEventsResponse Response Object
type ListPersonalRecentPushEventsResponse struct {

	// **参数解释：** 推送动态列表。
	Body           *[]PersonalPushEventDto `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListPersonalRecentPushEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPersonalRecentPushEventsResponse struct{}"
	}

	return strings.Join([]string{"ListPersonalRecentPushEventsResponse", string(data)}, " ")
}

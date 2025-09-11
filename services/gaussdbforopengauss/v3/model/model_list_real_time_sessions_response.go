package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRealTimeSessionsResponse Response Object
type ListRealTimeSessionsResponse struct {

	// **参数解释**： 数据库实例的实时会话列表。
	Sessions       *[]RealTimeSessionResult `json:"sessions,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListRealTimeSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRealTimeSessionsResponse struct{}"
	}

	return strings.Join([]string{"ListRealTimeSessionsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionsResponse Response Object
type ListSessionsResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 企业的会话列表。
	Items          *[]SessionInfo `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionsResponse struct{}"
	}

	return strings.Join([]string{"ListSessionsResponse", string(data)}, " ")
}

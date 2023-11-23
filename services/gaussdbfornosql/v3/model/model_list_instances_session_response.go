package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesSessionResponse Response Object
type ListInstancesSessionResponse struct {

	// 符合查询条件的总会话数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 实例会话详细信息列表。
	Sessions       *[]ListInstancesSessionRespondBodySessions `json:"sessions,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o ListInstancesSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesSessionResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesSessionResponse", string(data)}, " ")
}

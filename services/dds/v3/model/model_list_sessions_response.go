package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSessionsResponse struct {
	// 总记录数。

	TotalCount *int32 `json:"total_count,omitempty"`
	// 具体信息。

	Sessions       *[]QuerySessionResponse `json:"sessions,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionsResponse struct{}"
	}

	return strings.Join([]string{"ListSessionsResponse", string(data)}, " ")
}

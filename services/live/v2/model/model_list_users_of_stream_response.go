package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListUsersOfStreamResponse struct {

	// 观众趋势列表。
	DataList *[]V2UserData `json:"data_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListUsersOfStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersOfStreamResponse struct{}"
	}

	return strings.Join([]string{"ListUsersOfStreamResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcUserListResponse Response Object
type ListRtcUserListResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 查询结果限制
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 用户列表
	Users *[]RtcUser `json:"users,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcUserListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcUserListResponse struct{}"
	}

	return strings.Join([]string{"ListRtcUserListResponse", string(data)}, " ")
}

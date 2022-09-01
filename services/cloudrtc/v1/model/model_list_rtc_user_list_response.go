package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcUserListResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 查询结果限制
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 用户列表
	Users *[]RtcUser `json:"users,omitempty" xml:"users"`

	XRequestId     *string `json:"X-request-id,omitempty" xml:"X-request-id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcUserListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcUserListResponse struct{}"
	}

	return strings.Join([]string{"ListRtcUserListResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNoticesResponse struct {

	// 推荐公告列表
	Notices *[]Notice `json:"notices,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListNoticesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNoticesResponse struct{}"
	}

	return strings.Join([]string{"ListNoticesResponse", string(data)}, " ")
}

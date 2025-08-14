package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionsResponse Response Object
type ListSessionsResponse struct {

	// 用户登录会话列表
	SessionList    *[]UserSessionDto `json:"session_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionsResponse struct{}"
	}

	return strings.Join([]string{"ListSessionsResponse", string(data)}, " ")
}

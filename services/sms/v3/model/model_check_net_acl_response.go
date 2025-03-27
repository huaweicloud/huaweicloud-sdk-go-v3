package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckNetAclResponse Response Object
type CheckNetAclResponse struct {

	// 检查网卡安全组端口是否符合要求成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckNetAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNetAclResponse struct{}"
	}

	return strings.Join([]string{"CheckNetAclResponse", string(data)}, " ")
}

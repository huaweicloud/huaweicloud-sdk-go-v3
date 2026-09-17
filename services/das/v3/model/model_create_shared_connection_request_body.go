package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSharedConnectionRequestBody 创建共享链接请求体
type CreateSharedConnectionRequestBody struct {

	// 共享连接ID
	SharedConnId string `json:"shared_conn_id"`

	// 过期时间。格式为yyyy-MM-ddTHH:mm:ss.SSSZ
	ExpiredTime *string `json:"expired_time,omitempty"`

	// 用户列表
	Users []ShareConnUserInfo `json:"users"`
}

func (o CreateSharedConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSharedConnectionRequestBody", string(data)}, " ")
}

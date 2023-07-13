package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateShareConnectionsRequestBody struct {

	// 共享连接ID
	SharedConnId string `json:"shared_conn_id"`

	// 过期时间
	ExpiredTime *string `json:"expired_time,omitempty"`

	// 用户
	Users []ShareConnUserInfo `json:"users"`
}

func (o CreateShareConnectionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareConnectionsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateShareConnectionsRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CancelShareConnectionsRequestBody struct {

	// 共享链接ID
	SharedConnId string `json:"shared_conn_id"`

	// 用户
	Users []ShareConnUserInfo `json:"users"`
}

func (o CancelShareConnectionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelShareConnectionsRequestBody struct{}"
	}

	return strings.Join([]string{"CancelShareConnectionsRequestBody", string(data)}, " ")
}

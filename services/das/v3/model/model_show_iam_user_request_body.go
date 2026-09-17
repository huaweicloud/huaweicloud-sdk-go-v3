package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIamUserRequestBody 获取IAM用户信息请求体
type ShowIamUserRequestBody struct {

	// 账号ID
	UserIds *string `json:"user_ids,omitempty"`

	// 账号名称
	ConnectionId *string `json:"connection_id,omitempty"`
}

func (o ShowIamUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIamUserRequestBody struct{}"
	}

	return strings.Join([]string{"ShowIamUserRequestBody", string(data)}, " ")
}

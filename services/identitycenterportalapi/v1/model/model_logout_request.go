package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoutRequest Request Object
type LogoutRequest struct {

	// 创建令牌接口调用签发的访问令牌
	AccessToken string `json:"access-token"`
}

func (o LogoutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoutRequest struct{}"
	}

	return strings.Join([]string{"LogoutRequest", string(data)}, " ")
}

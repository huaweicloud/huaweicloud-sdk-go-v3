package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCssClusterReq 绑定CSS集群请求体
type CreateCssClusterReq struct {

	// 集群id
	CssId string `json:"css_id"`

	// 集群用户账号
	UserName string `json:"user_name"`

	// 集群用户密码，长度限制为[8,32]
	UserPwd string `json:"user_pwd"`
}

func (o CreateCssClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCssClusterReq struct{}"
	}

	return strings.Join([]string{"CreateCssClusterReq", string(data)}, " ")
}

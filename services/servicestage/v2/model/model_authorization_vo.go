package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 授权结构体。
type AuthorizationVo struct {

	// 授权名称。
	Name string `json:"name" xml:"name"`

	// 仓库类型。 取值范围：github、gitlab、gitee、bitbucket、devcloud。
	RepoType string `json:"repo_type" xml:"repo_type"`

	// 仓库地址。
	RepoHost string `json:"repo_host" xml:"repo_host"`

	// 仓库主页。
	RepoHome string `json:"repo_home" xml:"repo_home"`

	// 仓库用户名。
	RepoUser string `json:"repo_user" xml:"repo_user"`

	// 头像。
	Avartar string `json:"avartar" xml:"avartar"`

	// 授权方式。
	TokenType string `json:"token_type" xml:"token_type"`

	// 创建时间。
	CreateTime int64 `json:"create_time" xml:"create_time"`

	// 修改时间。
	UpdateTime int64 `json:"update_time" xml:"update_time"`

	// 状态。
	Status int32 `json:"status" xml:"status"`
}

func (o AuthorizationVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizationVo struct{}"
	}

	return strings.Join([]string{"AuthorizationVo", string(data)}, " ")
}

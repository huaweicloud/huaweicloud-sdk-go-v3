package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Repo 源码仓库信息。
type Repo struct {

	// 授权名称。
	AuthName *string `json:"auth_name,omitempty"`

	// 分支。
	Branch *string `json:"branch,omitempty"`

	// 命名空间，需填写命名空间的Base64编码。
	Namespace *string `json:"namespace,omitempty"`
}

func (o Repo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Repo struct{}"
	}

	return strings.Join([]string{"Repo", string(data)}, " ")
}

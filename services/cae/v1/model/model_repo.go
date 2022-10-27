package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Repo struct {

	// 授权名。
	AuthName *string `json:"auth_name,omitempty"`

	// 分支。
	Branch *string `json:"branch,omitempty"`

	// 命名空间。
	Namespace *string `json:"namespace,omitempty"`
}

func (o Repo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Repo struct{}"
	}

	return strings.Join([]string{"Repo", string(data)}, " ")
}

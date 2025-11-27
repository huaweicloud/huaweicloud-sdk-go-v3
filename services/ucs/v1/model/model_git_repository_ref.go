package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GitRepositoryRef struct {

	// 用于指定要检出的Git分支。如果未定义其他字段，则默认检出'master'分支。
	Branch *string `json:"branch,omitempty"`
}

func (o GitRepositoryRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GitRepositoryRef struct{}"
	}

	return strings.Join([]string{"GitRepositoryRef", string(data)}, " ")
}

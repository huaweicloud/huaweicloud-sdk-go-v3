package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoListHook struct {

	// hook列表
	Hooks *[]RepoHook `json:"hooks,omitempty"`
}

func (o RepoListHook) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoListHook struct{}"
	}

	return strings.Join([]string{"RepoListHook", string(data)}, " ")
}

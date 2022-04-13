package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRepoIdRequest struct {
	// 仓库组名

	GroupName string `json:"group_name"`
	// 仓库名

	RepositoryName string `json:"repository_name"`
}

func (o ShowRepoIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoIdRequest struct{}"
	}

	return strings.Join([]string{"ShowRepoIdRequest", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GetRepositoryByProjectIdRequest struct {
	// 仓库的uuid,用来指定需要查看的仓库

	RepositoryUuid string `json:"repository_uuid"`
}

func (o GetRepositoryByProjectIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetRepositoryByProjectIdRequest struct{}"
	}

	return strings.Join([]string{"GetRepositoryByProjectIdRequest", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectRepository struct {
	// 创建项目的UUID

	ProjectUuid *string `json:"projectUuid,omitempty"`
	// 创建仓库的UUID

	RepositoryUuid *string `json:"repositoryUuid,omitempty"`
}

func (o ProjectRepository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectRepository struct{}"
	}

	return strings.Join([]string{"ProjectRepository", string(data)}, " ")
}

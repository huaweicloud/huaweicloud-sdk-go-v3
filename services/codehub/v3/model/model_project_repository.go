package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectRepository struct {

	// 项目ID，获取方式请参见[获取项目ID](codehub_api_0014.xml)。
	ProjectUuid *string `json:"projectUuid,omitempty"`

	// 仓库UUID
	RepositoryUuid *string `json:"repositoryUuid,omitempty"`
}

func (o ProjectRepository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectRepository struct{}"
	}

	return strings.Join([]string{"ProjectRepository", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRepositoryNameExistRequest struct {

	// 项目的uuid
	ProjectUuid string `json:"project_uuid"`

	// 仓库名
	RepositoryName string `json:"repository_name"`
}

func (o ShowRepositoryNameExistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryNameExistRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryNameExistRequest", string(data)}, " ")
}

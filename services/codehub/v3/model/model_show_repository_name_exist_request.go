package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryNameExistRequest Request Object
type ShowRepositoryNameExistRequest struct {

	// 项目ID，获取方式请参见[获取项目ID](codehub_api_0014.xml)。
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockRepositoryRequest Request Object
type UnlockRepositoryRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 仓库的主键id
	RepositoryId int32 `json:"repository_id"`
}

func (o UnlockRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockRepositoryRequest struct{}"
	}

	return strings.Join([]string{"UnlockRepositoryRequest", string(data)}, " ")
}

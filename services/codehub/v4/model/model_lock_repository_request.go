package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockRepositoryRequest Request Object
type LockRepositoryRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 仓库的主键id
	RepositoryId int32 `json:"repository_id"`
}

func (o LockRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockRepositoryRequest struct{}"
	}

	return strings.Join([]string{"LockRepositoryRequest", string(data)}, " ")
}

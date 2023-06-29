package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryStatusRequest Request Object
type ListRepositoryStatusRequest struct {

	// 仓库的uuid,用来指定需要查看的仓库
	RepositoryUuid string `json:"repository_uuid"`
}

func (o ListRepositoryStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryStatusRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryStatusRequest", string(data)}, " ")
}

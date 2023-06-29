package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepositoryRequest Request Object
type DeleteRepositoryRequest struct {

	// 仓库uuid(由CreateRepository接口返回)，用来指定删除的仓库
	RepositoryUuid string `json:"repository_uuid"`
}

func (o DeleteRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepositoryRequest struct{}"
	}

	return strings.Join([]string{"DeleteRepositoryRequest", string(data)}, " ")
}

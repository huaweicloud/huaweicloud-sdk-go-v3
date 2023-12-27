package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryInfoRequest Request Object
type ShowRepositoryInfoRequest struct {

	// 仓库id
	RepoId string `json:"repo_id"`
}

func (o ShowRepositoryInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryInfoRequest", string(data)}, " ")
}

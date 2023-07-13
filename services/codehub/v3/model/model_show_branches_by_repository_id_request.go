package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBranchesByRepositoryIdRequest Request Object
type ShowBranchesByRepositoryIdRequest struct {

	// 仓库的主键id
	RepositoryId string `json:"repository_id"`
}

func (o ShowBranchesByRepositoryIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchesByRepositoryIdRequest struct{}"
	}

	return strings.Join([]string{"ShowBranchesByRepositoryIdRequest", string(data)}, " ")
}

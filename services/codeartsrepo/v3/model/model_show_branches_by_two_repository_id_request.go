package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBranchesByTwoRepositoryIdRequest Request Object
type ShowBranchesByTwoRepositoryIdRequest struct {

	// 仓库的主键id
	RepositoryId string `json:"repository_id"`
}

func (o ShowBranchesByTwoRepositoryIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchesByTwoRepositoryIdRequest struct{}"
	}

	return strings.Join([]string{"ShowBranchesByTwoRepositoryIdRequest", string(data)}, " ")
}

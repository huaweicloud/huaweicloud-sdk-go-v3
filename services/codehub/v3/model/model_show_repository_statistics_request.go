package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRepositoryStatisticsRequest struct {

	// 仓库的主键id
	RepositoryId int32 `json:"repository_id"`

	Body *ShowRepositoryStatisticsRequestBody `json:"body,omitempty"`
}

func (o ShowRepositoryStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryStatisticsRequest", string(data)}, " ")
}

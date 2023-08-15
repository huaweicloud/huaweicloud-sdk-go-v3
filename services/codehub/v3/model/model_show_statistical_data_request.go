package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticalDataRequest Request Object
type ShowStatisticalDataRequest struct {

	// 仓库id
	RepositoryUuid string `json:"repository_uuid"`
}

func (o ShowStatisticalDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticalDataRequest struct{}"
	}

	return strings.Join([]string{"ShowStatisticalDataRequest", string(data)}, " ")
}

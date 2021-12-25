package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowHasPipelineRequest struct {
	// 仓库id

	RepositoryUuid string `json:"repository_uuid"`
}

func (o ShowHasPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHasPipelineRequest struct{}"
	}

	return strings.Join([]string{"ShowHasPipelineRequest", string(data)}, " ")
}

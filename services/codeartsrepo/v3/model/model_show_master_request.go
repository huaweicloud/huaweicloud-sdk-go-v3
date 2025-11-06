package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMasterRequest Request Object
type ShowMasterRequest struct {

	// 仓库id
	RepositoryUuid string `json:"repository_uuid"`
}

func (o ShowMasterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMasterRequest struct{}"
	}

	return strings.Join([]string{"ShowMasterRequest", string(data)}, " ")
}

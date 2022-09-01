package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRepositoryByUuidRequest struct {

	// 仓库的uuid
	RepositoryUuid string `json:"repository_uuid" xml:"repository_uuid"`
}

func (o ShowRepositoryByUuidRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryByUuidRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryByUuidRequest", string(data)}, " ")
}

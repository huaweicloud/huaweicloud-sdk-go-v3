package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTenantRepositoryBody 仓库Id列表
type ExportTenantRepositoryBody struct {
	RepositoryIds *[]int32 `json:"repository_ids,omitempty"`
}

func (o ExportTenantRepositoryBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTenantRepositoryBody struct{}"
	}

	return strings.Join([]string{"ExportTenantRepositoryBody", string(data)}, " ")
}

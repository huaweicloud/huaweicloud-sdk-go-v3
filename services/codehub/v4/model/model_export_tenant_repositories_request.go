package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTenantRepositoriesRequest Request Object
type ExportTenantRepositoriesRequest struct {
	Body *ExportTenantRepositoryBody `json:"body,omitempty"`
}

func (o ExportTenantRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTenantRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ExportTenantRepositoriesRequest", string(data)}, " ")
}

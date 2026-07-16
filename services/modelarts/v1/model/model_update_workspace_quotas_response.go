package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkspaceQuotasResponse Response Object
type UpdateWorkspaceQuotasResponse struct {

	// 工作空间配额数据。
	Quotas         *[]WorkspaceQuotasUpdateResponse `json:"quotas,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o UpdateWorkspaceQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceQuotasResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceQuotasResponse", string(data)}, " ")
}

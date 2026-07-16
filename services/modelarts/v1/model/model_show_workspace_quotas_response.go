package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceQuotasResponse Response Object
type ShowWorkspaceQuotasResponse struct {

	// 工作空间配额数据。
	Quotas         *[]WorkspaceQuotasResponse `json:"quotas,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowWorkspaceQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceQuotasResponse", string(data)}, " ")
}

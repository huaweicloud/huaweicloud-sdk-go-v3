package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWorkspaceQuotasReq struct {

	// 工作空间配额数据。
	Quotas []UpdateWorkspaceQuotasReqQuotas `json:"quotas"`
}

func (o UpdateWorkspaceQuotasReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceQuotasReq struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceQuotasReq", string(data)}, " ")
}

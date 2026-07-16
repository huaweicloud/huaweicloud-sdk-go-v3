package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceQuotasRequest Request Object
type ShowWorkspaceQuotasRequest struct {

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId string `json:"workspace_id"`
}

func (o ShowWorkspaceQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceQuotasRequest", string(data)}, " ")
}

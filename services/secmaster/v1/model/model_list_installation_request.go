package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstallationRequest Request Object
type ListInstallationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o ListInstallationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstallationRequest struct{}"
	}

	return strings.Join([]string{"ListInstallationRequest", string(data)}, " ")
}

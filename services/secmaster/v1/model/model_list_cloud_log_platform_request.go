package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudLogPlatformRequest Request Object
type ListCloudLogPlatformRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o ListCloudLogPlatformRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudLogPlatformRequest struct{}"
	}

	return strings.Join([]string{"ListCloudLogPlatformRequest", string(data)}, " ")
}

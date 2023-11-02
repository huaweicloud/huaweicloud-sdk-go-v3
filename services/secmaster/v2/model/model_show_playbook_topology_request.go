package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookTopologyRequest Request Object
type ShowPlaybookTopologyRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 剧本实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowPlaybookTopologyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookTopologyRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookTopologyRequest", string(data)}, " ")
}

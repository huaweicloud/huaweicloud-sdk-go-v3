package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResRecallSetRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	// 资源id
	ResourceId string `json:"resource_id" xml:"resource_id"`

	// 使用类型： - UI，基于用户推荐物品 - UU，基于用户推荐用户 - II，基于物品推荐物品 - IU，基于物品推荐用户
	UseType string `json:"use_type" xml:"use_type"`
}

func (o ShowResRecallSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResRecallSetRequest struct{}"
	}

	return strings.Join([]string{"ShowResRecallSetRequest", string(data)}, " ")
}

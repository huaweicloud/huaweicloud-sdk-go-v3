package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResOnlineServiceDetailsRequest Request Object
type ListResOnlineServiceDetailsRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 资源id。
	ResourceId string `json:"resource_id"`

	// 服务类别： - SERVICE，在线服务
	Category string `json:"category"`
}

func (o ListResOnlineServiceDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResOnlineServiceDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListResOnlineServiceDetailsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResOnlineServiceDetailsRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	// 资源id。
	ResourceId string `json:"resource_id" xml:"resource_id"`

	// 服务类别： - SERVICE，在线服务
	Category string `json:"category" xml:"category"`
}

func (o ListResOnlineServiceDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResOnlineServiceDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListResOnlineServiceDetailsRequest", string(data)}, " ")
}

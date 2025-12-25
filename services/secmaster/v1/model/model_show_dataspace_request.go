package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataspaceRequest Request Object
type ShowDataspaceRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`
}

func (o ShowDataspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataspaceRequest struct{}"
	}

	return strings.Join([]string{"ShowDataspaceRequest", string(data)}, " ")
}

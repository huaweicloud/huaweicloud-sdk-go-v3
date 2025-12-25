package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataspaceRequest Request Object
type DeleteDataspaceRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`
}

func (o DeleteDataspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataspaceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDataspaceRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIsapComponentsRequest Request Object
type ListIsapComponentsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o ListIsapComponentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIsapComponentsRequest struct{}"
	}

	return strings.Join([]string{"ListIsapComponentsRequest", string(data)}, " ")
}

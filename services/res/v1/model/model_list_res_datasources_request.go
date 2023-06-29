package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResDatasourcesRequest Request Object
type ListResDatasourcesRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`
}

func (o ListResDatasourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResDatasourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResDatasourcesRequest", string(data)}, " ")
}

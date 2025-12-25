package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAdhocQueryRequest Request Object
type DeleteAdhocQueryRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 查询ID
	QueryId string `json:"query_id"`
}

func (o DeleteAdhocQueryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAdhocQueryRequest struct{}"
	}

	return strings.Join([]string{"DeleteAdhocQueryRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdhocResultRequest Request Object
type ShowAdhocResultRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 查询ID
	QueryId string `json:"query_id"`

	// 批次索引
	Batch int32 `json:"batch"`
}

func (o ShowAdhocResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdhocResultRequest struct{}"
	}

	return strings.Join([]string{"ShowAdhocResultRequest", string(data)}, " ")
}

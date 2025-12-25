package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableRequest Request Object
type ShowTableRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`
}

func (o ShowTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableRequest struct{}"
	}

	return strings.Join([]string{"ShowTableRequest", string(data)}, " ")
}

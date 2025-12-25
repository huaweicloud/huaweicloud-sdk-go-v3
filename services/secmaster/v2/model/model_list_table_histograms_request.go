package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableHistogramsRequest Request Object
type ListTableHistogramsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`

	Body *ListTableHistogramsRequestBody `json:"body,omitempty"`
}

func (o ListTableHistogramsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableHistogramsRequest struct{}"
	}

	return strings.Join([]string{"ListTableHistogramsRequest", string(data)}, " ")
}

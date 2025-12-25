package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchHistogramsRequest Request Object
type ListSearchHistogramsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *ListHistogramsRequestBody `json:"body,omitempty"`
}

func (o ListSearchHistogramsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchHistogramsRequest struct{}"
	}

	return strings.Join([]string{"ListSearchHistogramsRequest", string(data)}, " ")
}
